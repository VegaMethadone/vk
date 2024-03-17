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
		log.Fatalf("Unexpected error during getting films %v", err)
		//return nil
	}

	for res.Next() {
		newFilm := new(Film)

		res.Scan(&newFilm.Id, &newFilm.Name, &newFilm.Description, &newFilm.Enterdate, &newFilm.Rate, &newFilm.Score, &newFilm.Votes)

		acters, err := database.GetAllFilmActersDB(newFilm.Id)
		if err != nil {
			log.Fatalf("Unexpected error during getting films %v", err)
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
		return false
	}
	return result
}
