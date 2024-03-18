package film

import (
	"encoding/json"
	"log"
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
				log.Println(errorMessage, err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}
			if userData.Access < 1 {
				errorMessage := "Access denied"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(errorMessage))
				return
			}

			var acters []int
			err = json.Unmarshal([]byte(r.FormValue("acters")), &acters)
			if err != nil {
				errorMessage := "Error during parsing json"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}
			newDate := acter.ParseTime(r.FormValue("enterdate"))

			log.Println("Acters  data:", acters)
			for _, value := range acters {
				log.Println(value)
			}

			newFilm := &Film{
				Name:        r.FormValue("name"),
				Description: r.FormValue("description"),
				Enterdate:   newDate,
				Acters:      acters,
			}

			res := AddNewFilm(newFilm.Name, newFilm.Description, newFilm.Enterdate, newFilm.Acters)
			if res {
				log.Println("Successfully created film")
				w.WriteHeader(http.StatusCreated)
			} else {
				log.Println("Adding  failed", newFilm.Name)
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

func GetAllFilmsHandler(f *Film) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, err := authentication.CookieCheker(w, r)
			if err != nil {
				errorMessage := "Cookie is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}

			result := GetAllFilms()

			jsonData, err := json.Marshal(result)
			if err != nil {
				errorMessage := "Json is damaged"
				log.Println(errorMessage, err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}
			log.Println("Successfully parsed:", string(jsonData))
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

func ChangeFilmInfoHandler(f *Film) http.HandlerFunc {
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
				errorMessage := "Error during parsing json"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}
			if userData.Access < 1 {
				errorMessage := "Access denied"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(errorMessage))
				return
			}

			var acters []int
			err = json.Unmarshal([]byte(r.FormValue("acters")), &acters)
			if err != nil {
				errorMessage := "Json is damaged"
				log.Println(errorMessage, err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
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
				log.Println("Film info is changed:", newFilm.Id, newFilm.Name)
				w.WriteHeader(http.StatusOK)
			} else {
				log.Println("Changing film failed", newFilm.Id)
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

func GetFilmByFragmentHandler(f *Film) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, err := authentication.CookieCheker(w, r)
			if err != nil {
				errorMessage := "Cookie is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}

			result := GetFilmByFragment(r.FormValue("name"))

			jsonData, err := json.Marshal(result)
			if err != nil {
				errorMessage := "Json is damaged"
				log.Println(errorMessage, err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}

			log.Println("Json data:", string(jsonData))
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

func DeleteFilmHandler(f *Film) http.HandlerFunc {
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
				log.Println(errorMessage, err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}
			if userData.Access < 1 {
				errorMessage := "Access denied"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(errorMessage))
				return
			}

			targetID, _ := strconv.Atoi(r.FormValue("id"))
			log.Printf("Target ID: %d\n", targetID)

			result := DeleteFilm(targetID)
			if result {
				log.Println("Deleted  film:", targetID)
				w.WriteHeader(http.StatusOK)
			} else {
				log.Println("Delete film failed:", targetID)
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

func SortFilmsByRateHandler(f *Film) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, err := authentication.CookieCheker(w, r)
			if err != nil {
				errorMessage := "Cookie is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}

			result := GetAllFilms()
			result = SortFilmsByRate(result)

			jsonData, err := json.Marshal(result)
			if err != nil {
				errorMessage := "Json is damaged"
				log.Println(errorMessage, err)
				w.WriteHeader(http.StatusBadRequest)
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

func SortFilmsByNameHandler(f *Film) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, err := authentication.CookieCheker(w, r)
			if err != nil {
				errorMessage := "Cookie is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}

			result := GetAllFilms()
			result = SortFilmsByName(result)

			jsonData, err := json.Marshal(result)
			if err != nil {
				errorMessage := "Json is damaged"
				log.Println(errorMessage, err)
				w.WriteHeader(http.StatusBadRequest)
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

func SortFilmsByDateHandler(f *Film) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, err := authentication.CookieCheker(w, r)
			if err != nil {
				errorMessage := "Cookie is damaged"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}

			result := GetAllFilms()
			result = SortFilmsByDate(result)

			jsonData, err := json.Marshal(result)
			if err != nil {
				errorMessage := "Json is damaged"
				log.Println(errorMessage, err)
				w.WriteHeader(http.StatusBadRequest)
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
