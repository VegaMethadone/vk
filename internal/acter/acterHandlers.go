package acter

import (
	"encoding/json"
	"log"
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
				errorMessage := "Cookie is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}
			userData := new(authentication.User)
			err = json.Unmarshal(c, userData)
			if err != nil {
				errorMessage := "Json is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}
			if userData.Access < 1 {
				errorMessage := "Access denied"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
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
				errorMessage := "Server error during adding new acter"
				log.Println(errorMessage, newActer.Name, newActer.Name, newActer.DateOfBirth, newActer.Films)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}

			if result {
				log.Println("Added new acter")
				w.WriteHeader(http.StatusOK)
			} else {
				log.Println("Error adding new acter")
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			errorMessage := "NotAllowed"
			log.Println(errorMessage)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(errorMessage))
			return
		}
	}
}

func ChangeActerInfoHandler(a *Acter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			c, err := authentication.CookieCheker(w, r)
			if err != nil {
				errorMessage := "Cookie is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}
			userData := new(authentication.User)
			err = json.Unmarshal(c, userData)
			if err != nil {
				errorMessage := "Json is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}
			if userData.Access < 1 {
				errorMessage := "NotAllowed"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
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
				log.Println("Successfully changed acter")
				w.WriteHeader(http.StatusAccepted)
			} else {
				log.Println("Error change  info  acter")
				w.WriteHeader(http.StatusInternalServerError)
			}

		} else {
			errorMessage := "NotAllowed"
			log.Println(errorMessage)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(errorMessage))
			return
		}
	}
}

func DeleteActerInfoHandler(a *Acter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			c, err := authentication.CookieCheker(w, r)
			if err != nil {
				errorMessage := "Cookie is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}
			userData := new(authentication.User)
			err = json.Unmarshal(c, userData)
			if err != nil {
				errorMessage := "Json is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}
			if userData.Access < 1 {
				errorMessage := "NotAllowed"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}

			acterId, _ := strconv.Atoi(r.FormValue("id"))

			result := DeleteActer(acterId)
			if result {
				log.Println("Acter Id:", acterId, "Deleted")
				w.WriteHeader(http.StatusOK)
			} else {
				log.Println("Acter Id:", acterId, "Error  delete")
				w.WriteHeader(http.StatusInternalServerError)
			}

		} else {
			errorMessage := "NotAllowed"
			log.Println(errorMessage)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(errorMessage))
			return
		}
	}
}

func GetAllActersHandler(a *Acter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, err := authentication.CookieCheker(w, r)
			if err != nil {
				errorMessage := "Cookie is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(errorMessage))
				return
			}

			result := GetAllActersList()
			log.Println("RESULT: ", result)

			jsonData, err := json.Marshal(result)
			log.Println("JSON DATA:", string(jsonData))
			if err != nil {
				errorMessage := "Json is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		} else {
			errorMessage := "NotAllowed"
			log.Println(errorMessage)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(errorMessage))
			return
		}
	}
}
