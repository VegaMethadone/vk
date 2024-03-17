package main

import (
	"fmt"
	"net/http"

	"vk/internal/acter"
	"vk/internal/authentication"
	"vk/internal/film"
)

func main() {

	//used structers
	newUser := authentication.User{}
	newActer := acter.Acter{}
	newFilm := film.Film{}

	//main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет от сервера!")
	})

	// user & authentication
	http.HandleFunc("/UserRegistration", authentication.RegistrationHandler(&newUser))
	http.HandleFunc("/EnterUser", authentication.EnterUserHandler(&newUser))
	http.HandleFunc("/ChangeUserData", authentication.ChangeUserDataHandler(&newUser))

	// acters
	http.HandleFunc("/AddActer", acter.AddNewActerHandler(&newActer))
	http.HandleFunc("/ChangeActerInfo", acter.ChangeActerInfoHandler(&newActer))
	http.HandleFunc("/DeleteActerInfo", acter.DeleteActerInfoHandler(&newActer))
	http.HandleFunc("/GetAllActers", acter.GetAllActersHandler(&newActer))

	// films
	http.HandleFunc("/AddFilm", film.AddNewFilmHandler(&newFilm))
	http.HandleFunc("/GetAllFilms", film.GetAllFilmsHandler(&newFilm))

	//Server
	fmt.Println("Server is working  at http://127.0.0.1:8080")
	//http://localhost:8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Start server error")
	}

}
