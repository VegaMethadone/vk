package acter

import (
	"errors"
	"log"
	"time"

	"vk/internal/database"
)

func (a *Acter) ChangeName(name string) error {

	a.Name = name
	return nil
}

func (a *Acter) ChangeSex(sex string) error {
	if sex != "male" && sex != "female" && sex != "" {
		return errors.New("uncorrect input")
	}
	a.Sex = sex
	return nil
}

func (a *Acter) ChangeDate(dateOfBirth string) error {
	maket := "2006-01-02"
	parsedDate, err := time.Parse(maket, dateOfBirth)
	if err != nil {
		log.Println(err)
		return err
	}

	if parsedDate.After(time.Now()) {
		return errors.New("uncorrect input")
	}
	a.DateOfBirth = parsedDate
	return nil
}

func GetAllActersList() []Acter {
	arrActers := []Acter{}

	res, err := database.GetAllActersDB()
	if err != nil {
		log.Fatalf("Did not get acters from DB %v\n", err)
	}

	for res.Next() {
		newActer := new(Acter)
		res.Scan(&newActer.Id, &newActer.Name, &newActer.Sex, &newActer.DateOfBirth)

		films, err := database.GetALLActerFilmsDB(newActer.Id)
		if err != nil {
			log.Fatalf("Did not get acter films from DB %v\n", err)
		}

		for films.Next() {
			var filmId int
			films.Scan(&filmId)
			newActer.Films = append(newActer.Films, filmId)
		}
		arrActers = append(arrActers, *newActer)
	}
	return arrActers
}
