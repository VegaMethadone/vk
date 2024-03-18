package authentication

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func RegistrationHandler(u *User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			accessStr := r.FormValue("access")
			tmpAccess, err := strconv.Atoi(accessStr)
			if err != nil {
				errorMessage := "Invalid access value"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}

			newUser := &User{
				Id:       0,
				Login:    r.FormValue("login"),
				Password: r.FormValue("password"),
				Access:   tmpAccess,
			}

			result := newUser.CreateUser(newUser.Login, newUser.Password, newUser.Access)
			if result {
				message := "Successful registration"
				log.Println(message, newUser.Login, newUser.Access)
				w.WriteHeader(http.StatusCreated)
				fmt.Fprintf(w, "Successful registration\n")
			} else {
				errorMessage := "Registration  failed"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
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

func EnterUserHandler(u *User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			findUser := &User{
				Login:    r.FormValue("login"),
				Password: r.FormValue("password"),
			}

			result, foundUser := findUser.FindUser(findUser.Login, findUser.Password)
			if result {
				userJSON, err := json.Marshal(foundUser)
				if err != nil {
					errorMessage := "Json Marshal Error  during authorization"
					log.Println(errorMessage)
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(errorMessage))
					return
				}

				token := base64Encode(userJSON)
				CookieSeter(w, r, token)

				fmt.Fprintf(w, "Successfuly found\n")
			} else {
				errorMessage := "Faild authorization"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
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

func ChangeUserDataHandler(u *User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			c, err := r.Cookie("session_token")
			if err != nil {
				errorMessage := "No cookie"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errorMessage))
				return
			}

			oldJson, _ := base64Decode(c.Value)
			oldUser := new(User)
			json.Unmarshal(oldJson, oldUser)

			newUser := &User{
				Id:       oldUser.Id,
				Login:    r.FormValue("login"),
				Password: r.FormValue("password"),
				Access:   oldUser.Access,
			}

			fmt.Printf("newData: %v\n", *newUser)

			result, _ := newUser.ChangeUserData(newUser.Id, newUser.Access, newUser.Login, newUser.Password)
			if !result {
				errorMessage := "Server error during  changing user data"
				log.Println(errorMessage)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return

			} else {
				newJson, err := json.Marshal(newUser)
				if err != nil {
					errorMessage := "Json Marshal Error  during authorization"
					log.Println(errorMessage)
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(errorMessage))
					return
				}

				token := base64Encode(newJson)
				log.Printf("New token: %s\n", token)

				CookieSeter(w, r, token)

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
