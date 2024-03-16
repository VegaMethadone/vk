package database

import (
	"fmt"
	"log"
	"testing"
	"time"
)

type testActer struct {
	id          int
	name        string
	sex         string
	dateOfBirth time.Time
	films       []int
}

func TestAddNewActerDB(t *testing.T) {
	acterOne := &testActer{
		name:        "Ryan Gosling",
		sex:         "male",
		dateOfBirth: time.Date(1980, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	acterTwo := &testActer{
		name:        "Mads Mikkelsen",
		sex:         "male",
		dateOfBirth: time.Date(1965, time.September, 22, 0, 0, 0, 0, time.UTC),
	}

	arr := []testActer{*acterOne, *acterTwo}

	for key, value := range arr {
		res, err := AddNewActerDB(value.name, value.sex, value.dateOfBirth)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !res {
			t.Errorf("Acter %d was not added", key)
		}
		fmt.Println(res)
	}

}

func TestFindActerDB(t *testing.T) {
	acterOne := &testActer{
		name:        "Ryan Gosling",
		sex:         "male",
		dateOfBirth: time.Date(1980, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	acterTwo := &testActer{
		name:        "Mads Mikkelsen",
		sex:         "male",
		dateOfBirth: time.Date(1965, time.September, 22, 0, 0, 0, 0, time.UTC),
	}
	arr := []testActer{*acterOne, *acterTwo}

	for key, value := range arr {
		flag, res, err := FindActerDB(value.name)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !flag {
			t.Errorf("Acter %d was not found", key)
		}

		for res.Next() {
			var id int
			var name string
			var sex string
			var dateOfBirth time.Time

			_ = res.Scan(&id, &name, &sex, &dateOfBirth)
			fmt.Println("EXPECTED:", value.name, value.sex, value.dateOfBirth, "\nGOT:", name, sex, dateOfBirth)
		}

	}
}

func TestChangeActerNameDB(t *testing.T) {
	acterOne := &testActer{
		name:        "Ryan Gosling",
		sex:         "male",
		dateOfBirth: time.Date(1980, time.November, 12, 0, 0, 0, 0, time.UTC),
	}

	flag, err := ChangeActerNameDB(acterOne.id, "Vasya Pupkin")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !flag {
		t.Errorf("Acter %d was not found", 1)
	}

	flag, res, err := FindActerDB("Vasya Pupkin")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !flag {
		t.Errorf("Acter %d was not found", 1)
	}

	for res.Next() {
		var id int
		var name string
		var sex string
		var dateOfBirth time.Time

		_ = res.Scan(&id, &name, &sex, &dateOfBirth)
		fmt.Println("OLD:", acterOne.name, acterOne.sex, acterOne.dateOfBirth, "\nNEW:", name, sex, dateOfBirth)
	}

}

func TestChangeActerSexDB(t *testing.T) {
	acterOne := &testActer{
		id:          2,
		name:        "Mads Mikkelsen",
		sex:         "male",
		dateOfBirth: time.Date(1965, time.September, 22, 0, 0, 0, 0, time.UTC),
	}

	flag, err := ChangeActerSexDB(acterOne.id, "female")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !flag {
		t.Errorf("Acter %d was not found", 1)
	}

	flag, res, err := FindActerDB("Mads Mikkelsen")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !flag {
		t.Errorf("Acter %d was not found", 2)
	}

	for res.Next() {
		var id int
		var name string
		var sex string
		var dateOfBirth time.Time

		_ = res.Scan(&id, &name, &sex, &dateOfBirth)
		fmt.Println("OLD:", acterOne.name, acterOne.sex, acterOne.dateOfBirth, "\nNEW:", name, sex, dateOfBirth)
	}
}

func TestChangeActerDateOfBirthDB(t *testing.T) {
	acterOne := &testActer{
		id:          2,
		name:        "Mads Mikkelsen",
		sex:         "male",
		dateOfBirth: time.Date(1965, time.September, 22, 0, 0, 0, 0, time.UTC),
	}
	flag, err := ChangeActerDateOfBirthDB(acterOne.id, time.Date(2024, time.March, 14, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !flag {
		t.Errorf("Acter %d was not found", 1)
	}

	flag, res, err := FindActerDB("Mads Mikkelsen")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !flag {
		t.Errorf("Acter %d was not found", 2)
	}

	for res.Next() {
		var id int
		var name string
		var sex string
		var dateOfBirth time.Time

		_ = res.Scan(&id, &name, &sex, &dateOfBirth)
		fmt.Println("OLD:", acterOne.name, acterOne.sex, acterOne.dateOfBirth, "\nNEW:", name, sex, dateOfBirth)
	}
}

func TestChangeActerAllDB(t *testing.T) {
	acterOne := &testActer{
		id:          2,
		name:        "Mads Mikkelsen",
		sex:         "male",
		dateOfBirth: time.Date(1965, time.September, 22, 0, 0, 0, 0, time.UTC),
	}

	flag, err := ChangeActerAllDB(acterOne.id, "Daniel Day-Lewis", "male", time.Date(1957, time.April, 29, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !flag {
		t.Errorf("Acter %d was not found", 2)
	}

	flag, res, err := FindActerDB("Daniel Day-Lewis")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !flag {
		t.Errorf("Acter %d was not found", 2)
	}

	for res.Next() {
		var id int
		var name string
		var sex string
		var dateOfBirth time.Time

		_ = res.Scan(&id, &name, &sex, &dateOfBirth)
		fmt.Println("OLD:", acterOne.name, acterOne.sex, acterOne.dateOfBirth, "\nNEW:", name, sex, dateOfBirth)
	}
}

func TestDeleteActerInfoDB(t *testing.T) {
	acterOne := &testActer{
		name:        "Ryan Gosling",
		sex:         "male",
		dateOfBirth: time.Date(1980, time.November, 12, 0, 0, 0, 0, time.UTC),
	}
	flag, err := DeleteActerInfoDB(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !flag {
		t.Errorf("Acter %d was not found", 2)
	}

	flag, res, err := FindActerDB("Ryan Gosling")
	if err != nil {
		t.Skip(err)
	}
	if !flag {
		t.Skip(err)
	}

	for res.Next() {
		var id int
		var name string
		var sex string
		var dateOfBirth time.Time

		_ = res.Scan(&id, &name, &sex, &dateOfBirth)
		fmt.Println("OLD:", acterOne.name, acterOne.sex, acterOne.dateOfBirth, "\nNEW:", name, sex, dateOfBirth)
	}

}

func TestGetAllActersDB(t *testing.T) {
	allActers := []testActer{}

	res, err := GetAllActersDB()
	if err != nil {
		log.Fatalf("Unexpected error ocured %v", err)
	}

	for res.Next() {
		newActer := new(testActer)
		res.Scan(&newActer.id, &newActer.name, &newActer.sex, &newActer.dateOfBirth)

		films, err := GetALLActerFilmsDB(newActer.id)
		if err != nil {
			log.Fatalf("Unexpected error %v", err)
		}

		for films.Next() {
			var tmpFilm int
			films.Scan(&tmpFilm)
			newActer.films = append(newActer.films, tmpFilm)
		}

		allActers = append(allActers, *newActer)
	}

	for _, value := range allActers {
		log.Printf("%v\n", value)
	}
}
