package acter

import (
	"time"
)

type Acter struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Sex         string    `json:"sex"`
	DateOfBirth time.Time `json:"date"`
	Films       []int     `json:"films"`
}

/*
type Acter interface {
	ChangeName(string) error
	ChangeSex(string) error
	ChangeDate(string) error
	DeleteActer() error
	AddActer(string, string, time.Time) error
	ChangeActerInfoAll()
	GetAllActersList() []acter
}
*/
