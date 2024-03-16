package acter

import (
	"fmt"
	"net/http"
)

func ChangeActerNameHandler(a *Acter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newName := r.FormValue("name")
		a.ChangeName(newName)
		fmt.Fprint(w, "Name is changed %s", newName)
	}
}

func ChangeActerSex(a *Acter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newSex := r.FormValue("sex")
		a.ChangeSex(newSex)
		fmt.Fprintf(w, "Sex is changed %s", newSex)
	}
}

func ChangeActerDate(a *Acter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newDate := r.FormValue("date")
		a.ChangeDate(newDate)
		fmt.Fprintf(w, "Birthday date is changed %s", newDate)
	}
}
