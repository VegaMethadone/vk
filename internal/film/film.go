package film

import (
	"log"
	"strings"
	"time"

	"vk/internal/database"
)

func GetAllFilms() []Film {
	arrFilms := []Film{}

	res, err := database.GetAllFilmsDB()
	if err != nil {
		log.Printf("Unexpected error during getting films %v", err)
		return nil
	}

	for res.Next() {
		newFilm := new(Film)

		res.Scan(&newFilm.Id, &newFilm.Name, &newFilm.Description, &newFilm.Enterdate, &newFilm.Rate, &newFilm.Score, &newFilm.Votes)

		acters, err := database.GetAllFilmActersDB(newFilm.Id)
		if err != nil {
			log.Printf("Unexpected error during getting films %v", err)
			return nil
		}

		for acters.Next() {
			var acterId int
			acters.Scan(&acterId)
			newFilm.Acters = append(newFilm.Acters, acterId)
		}
		arrFilms = append(arrFilms, *newFilm)
	}
	return arrFilms
}

func AddNewFilm(name, description string, enterdate time.Time, acters []int) bool {
	name = strings.ToLower(name)

	result, err := database.AddNewFilmDB(name, description, enterdate, acters)
	if err != nil {
		log.Println("Error occured during AddNewFilmDB:", err)
		return false
	}
	return result
}

func ChangeFilmInfo(id int, name, description string, enterdate time.Time, score, votes int, acters []int) bool {
	newName := strings.ToLower(name)
	result, err := database.ChangeFilmInfoDB(id, newName, description, enterdate, score, votes, acters)
	if err != nil {
		log.Println("Error occured during ChangeFilmInfoDB:", err)
		return false
	}

	return result
}

func GetFilmByFragment(fragment string) []Film {
	fragment = strings.ToLower(fragment)

	_, res, err := database.FindFilmByFragmentDB(fragment)
	if err != nil {
		log.Printf("Not found  %v\n", err)
		return nil
	}

	filmArr := []Film{}
	for res.Next() {
		film := new(Film)
		res.Scan(&film.Id, &film.Name, &film.Description, &film.Enterdate, &film.Rate, &film.Score, &film.Votes)

		acters, err := database.GetAllFilmActersDB(film.Id)
		if err != nil {
			log.Println("Error occurred GetAllFilmActersDB:", err)
		}
		for acters.Next() {
			var acterId int
			acters.Scan(&acterId)
			film.Acters = append(film.Acters, acterId)
		}
		filmArr = append(filmArr, *film)
	}
	return filmArr
}

func DeleteFilm(id int) bool {
	database.DeleteFilmActersDB(id)

	res, err := database.DeleteFilmDB(id)
	if err != nil {
		log.Println("Error occurred DeleteFilmDB:", err)
		return false
	}
	return res
}
