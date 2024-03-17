package film

import (
	"log"
	"testing"
)

func TestSortFilmsByRate(t *testing.T) {
	result := GetAllFilms()
	result = SortFilmsByRate(result)

	for key, value := range result {
		log.Printf("Index:  %d, Name: %s,  Rate: %f\n", key, value.Name, value.Rate)
	}
}

func TestSortFilmsByName(t *testing.T) {
	result := GetAllFilms()
	result = SortFilmsByName(result)

	for key, value := range result {
		log.Printf("Index:  %d, Name: %s\n", key, value.Name)
	}
}

func TestSortFilmsByDate(t *testing.T) {
	result := GetAllFilms()
	result = SortFilmsByDate(result)

	for key, value := range result {
		log.Printf("Index:  %d, Name: %s, Date: %v\n", key, value.Name, value.Enterdate)
	}
}
