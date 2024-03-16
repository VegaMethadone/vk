package film

import (
	"time"
)

type film struct {
	id          int
	name        string // не более 150 символов
	description string //  не более 1000 символов
	enterdate   time.Time
	rate        float64
	score       int
	acters      []int
}

type Film interface {
	AddNewFilm(string, string, time.Time, []int)
	FindFilm()
	ChangeFilmInfo()
	DeleteFilm()
	GetAllFilms() []film
}
