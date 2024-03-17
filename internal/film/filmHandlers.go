package film

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"vk/internal/acter"
	"vk/internal/authentication"
)

func AddNewFilmHandler(f *Film) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			c, err := authentication.CookieCheker(w, r)
			if err != nil {
				http.Error(w, "Cookie is damaged", http.StatusBadRequest)
				return
			}
			userData := new(authentication.User)
			err = json.Unmarshal(c, userData)
			if err != nil {
				http.Error(w, "Cookie is damaged", http.StatusBadRequest)
				return
			}
			if userData.Access < 1 {
				http.Error(w, "Access denied", http.StatusUnauthorized)
				return
			}

			var acters []int
			err = json.Unmarshal([]byte(r.FormValue("acters")), &acters)
			if err != nil {
				http.Error(w, "Error during parsing json", http.StatusInternalServerError)
				return
			}
			newDate := acter.ParseTime(r.FormValue("enterdate"))
			//COMMENT
			fmt.Println("Acters  data:", acters)
			for _, value := range acters {
				fmt.Println(value)
			}
			//COMMENT

			newFilm := &Film{
				Name:        r.FormValue("name"),
				Description: r.FormValue("description"),
				Enterdate:   newDate,
				Acters:      acters,
			}

			res := AddNewFilm(newFilm.Name, newFilm.Description, newFilm.Enterdate, newFilm.Acters)
			if res {
				w.WriteHeader(http.StatusCreated)
			} else {
				http.Error(w, "Adding  failed", http.StatusInternalServerError)
			}

		} else {
			http.Error(w, "Not allowed", http.StatusBadRequest)
		}
	}
}

func GetAllFilmsHandler(f *Film) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, err := authentication.CookieCheker(w, r)
			if err != nil {
				http.Error(w, "Cookie is damaged", http.StatusBadRequest)
				return
			}

			result := GetAllFilms()

			jsonData, err := json.Marshal(result)
			if err != nil {
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		} else {
			http.Error(w, "Not allowed", http.StatusBadRequest)
		}
	}
}

func ChangeFilmInfoHandler(f *Film) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			c, err := authentication.CookieCheker(w, r)
			if err != nil {
				http.Error(w, "Cookie is damaged", http.StatusBadRequest)
				return
			}
			userData := new(authentication.User)
			err = json.Unmarshal(c, userData)
			if err != nil {
				http.Error(w, "Cookie is damaged", http.StatusBadRequest)
				return
			}
			if userData.Access < 1 {
				http.Error(w, "Access denied", http.StatusUnauthorized)
				return
			}

			var acters []int
			err = json.Unmarshal([]byte(r.FormValue("acters")), &acters)
			if err != nil {
				http.Error(w, "Error during parsing json", http.StatusInternalServerError)
				return
			}
			oldId, _ := strconv.Atoi(r.FormValue("id"))
			newDate := acter.ParseTime(r.FormValue("enterdate"))
			newScore, _ := strconv.Atoi(r.FormValue("score"))
			newVotes, _ := strconv.Atoi(r.FormValue("votes"))

			newFilm := &Film{
				Id:          oldId,
				Name:        r.FormValue("name"),
				Description: r.FormValue("description"),
				Enterdate:   newDate,
				Score:       newScore,
				Votes:       newVotes,
				Acters:      acters,
			}

			result := ChangeFilmInfo(newFilm.Id, newFilm.Name, newFilm.Description, newFilm.Enterdate, newFilm.Score, newFilm.Votes, newFilm.Acters)

			if result {
				w.WriteHeader(http.StatusOK)
			} else {
				http.Error(w, "Adding  failed", http.StatusInternalServerError)
			}

		} else {
			http.Error(w, "Not allowed", http.StatusBadRequest)
		}
	}
}
