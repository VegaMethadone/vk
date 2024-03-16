package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func RegistrationHandler(u *User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			accessStr := r.FormValue("access")
			tmpAccess, err := strconv.Atoi(accessStr)
			if err != nil {
				http.Error(w, "Invalid access value", http.StatusBadRequest)
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
				w.WriteHeader(http.StatusCreated)
				fmt.Fprintf(w, "Successful registration\n")
			} else {
				http.Error(w, "Registration failed", http.StatusInternalServerError)
			}

		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
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
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
				//fix
				token := base64Encode(userJSON)
				CookieSeter(w, r, token)

				fmt.Fprintf(w, "Successfuly found\n")
			} else {
				http.Error(w, "failed", http.StatusNotFound)
				return
			}

		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
	}

}

func ChangeUserDataHandler(u *User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			c, err := r.Cookie("session_token")
			if err != nil {
				http.Error(w, "Not log in", http.StatusBadRequest)
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
				http.Error(w, "Server  error\n", http.StatusInternalServerError)
				return
			} else {
				newJson, err := json.Marshal(newUser)
				if err != nil {
					http.Error(w, "Error during Marshal JSON\n", http.StatusInternalServerError)
					return
				}
				fmt.Printf("New JSON:  %v\n", string(newJson))

				token := base64Encode(newJson)
				fmt.Printf("New token: %s\n", token)

				CookieSeter(w, r, token)

			}

		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}

}
