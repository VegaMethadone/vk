package film

import (
	"log"
	"time"

	"vk/internal/database"
)

func GetAllFilms() []film {
	arrFilms := []film{}

	res, err := database.GetAllFilmsDB()
	if err != nil {
		log.Fatalf("Unexpected error during getting films %v", err)
		//return nil
	}

	for res.Next() {
		newFilm := new(film)

		res.Scan(&newFilm.id, &newFilm.name, &newFilm.description, &newFilm.enterdate, &newFilm.rate, &newFilm.score)

		acters, err := database.GetAllFilmActersDB(newFilm.id)
		if err != nil {
			log.Fatalf("Unexpected error during getting films %v", err)
			//return nil
		}

		for acters.Next() {
			var acterId int
			acters.Scan(&acterId)
			newFilm.acters = append(newFilm.acters, acterId)
		}
		arrFilms = append(arrFilms, *newFilm)
	}
	return arrFilms
}

func AddNewFilm(name, description string, enterdate time.Time, acters []int) (bool, error) {
	result, err := database.AddNewFilmDB(name, description, enterdate, acters)
	if err != nil {
		log.Fatalf("Unexpected error during adding films %v", err)
		return result, err
	}
	return result, nil
}

func ChangeFilmInfo(name, newName, newDescription string, newEnterdate time.Time, newScore int, changeName, changeDescription, changeEnterdate, changeScore bool) (bool, error) {
	result, err := database.ChangeFilmInfoDB(name, newName, newDescription, newEnterdate, newScore, changeName, changeDescription, changeEnterdate, changeScore)
	if err != nil {
		log.Fatalf("Unexpected error during changing info %v", err)
		return result, err
	}

	return result, nil
}
