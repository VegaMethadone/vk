package acter

import (
	"fmt"
	"log"
	"strings"
	"time"

	"vk/internal/database"
)

func ParseTime(str string) time.Time {
	layout := "2006-01-02"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error during parsing  time:", err)
		return time.Time{}
	}
	return t
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

func AddNewActer(name, sex string, dateOfBirth time.Time) bool {
	name = strings.ToLower(name)
	sex = strings.ToLower(sex)

	result, err := database.AddNewActerDB(name, sex, dateOfBirth)
	if err != nil {
		log.Fatalf("Error during adding new actor  %v", err)
		return false
	}

	return result
}

func ChangeActerInfo(id int, name, sex string, dateOfBirth time.Time) bool {
	name = strings.ToLower(name)
	sex = strings.ToLower(sex)
	if dateOfBirth.After(time.Now()) {
		log.Fatalf("Incorrect date\n")
		return false
	}

	_, err := database.FindActerByIdDB(id)
	if err != nil {
		log.Fatalf("Error during finding acter  %v", err)
		return false
	}

	result, err := database.ChangeActerAllDB(id, name, sex, dateOfBirth)
	if err != nil {
		log.Fatalf("Database error %v", err)
		return false
	}
	return result
}

func DeleteActer(id int) bool {
	err := database.DeleteAllActerFilmsDB(id)
	if err != nil {
		return false
	}

	result, err := database.DeleteActerInfoDB(id)
	if err != nil {
		log.Fatalf("Database error")
		return false
	}
	return result
}
