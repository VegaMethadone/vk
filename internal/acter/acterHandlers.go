package acter

import (
	"encoding/json"
	"net/http"
	"strconv"

	"vk/internal/authentication"
	"vk/internal/database"
)

func AddNewActerHandler(a *Acter) http.HandlerFunc {
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

			parsedTime := ParseTime(r.FormValue("date"))
			newActer := &Acter{
				Name:        r.FormValue("name"),
				Sex:         r.FormValue("sex"),
				DateOfBirth: parsedTime,
			}

			result, err := database.AddNewActerDB(newActer.Name, newActer.Sex, newActer.DateOfBirth)
			if err != nil {
				http.Error(w, "Server  error", http.StatusInternalServerError)
				return
			}

			if result {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			http.Error(w, "Not allowed", http.StatusBadRequest)
		}
	}
}

func ChangeActerInfoHandler(a *Acter) http.HandlerFunc {
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

			parsedTime := ParseTime(r.FormValue("date"))
			parsedId, _ := strconv.Atoi(r.FormValue("id"))
			updateActer := &Acter{
				Id:          parsedId,
				Name:        r.FormValue("name"),
				Sex:         r.FormValue("sex"),
				DateOfBirth: parsedTime,
			}

			result := ChangeActerInfo(updateActer.Id, updateActer.Name, updateActer.Sex, updateActer.DateOfBirth)
			if result {
				w.WriteHeader(http.StatusAccepted)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

		} else {
			http.Error(w, "Not allowed", http.StatusBadRequest)
		}
	}
}

func DeleteActerInfoHandler(a *Acter) http.HandlerFunc {
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

			acterId, _ := strconv.Atoi(r.FormValue("id"))

			result := DeleteActer(acterId)
			if result {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

		} else {
			http.Error(w, "Not allowed", http.StatusBadRequest)
		}
	}
}

func GetAllActersHandler(a *Acter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, err := authentication.CookieCheker(w, r)
			if err != nil {
				http.Error(w, "Cookie is damaged", http.StatusBadRequest)
				return
			}

			result := GetAllActersList()

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
