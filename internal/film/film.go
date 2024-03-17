package film

import (
	"log"

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
