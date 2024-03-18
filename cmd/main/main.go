package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"vk/internal/acter"
	"vk/internal/authentication"
	"vk/internal/film"
)

func main() {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	newDirPath := filepath.Join(filepath.Dir(filepath.Dir(currentDir)), "logs")

	newDirPath += "data.log"
	file, err := os.OpenFile(newDirPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Println("Did not create/find log file")
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("Server is strarting")

	//used structers
	newUser := authentication.User{}
	newActer := acter.Acter{}
	newFilm := film.Film{}

	//main page
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is working!")
	})

	// user & authentication
	http.HandleFunc("/user/UserRegistration", authentication.RegistrationHandler(&newUser))
	http.HandleFunc("/user/EnterUser", authentication.EnterUserHandler(&newUser))
	http.HandleFunc("/user/ChangeUserData", authentication.ChangeUserDataHandler(&newUser))

	// acters
	http.HandleFunc("/acter/AddActer", acter.AddNewActerHandler(&newActer))
	http.HandleFunc("/acter/ChangeActerInfo", acter.ChangeActerInfoHandler(&newActer))
	http.HandleFunc("/acter/DeleteActerInfo", acter.DeleteActerInfoHandler(&newActer))
	http.HandleFunc("/acter/GetAllActers", acter.GetAllActersHandler(&newActer))

	// films
	http.HandleFunc("/films/AddFilm", film.AddNewFilmHandler(&newFilm))
	http.HandleFunc("/films/GetAllFilms", film.GetAllFilmsHandler(&newFilm))
	http.HandleFunc("/films/ChangeFilmInfo", film.ChangeFilmInfoHandler(&newFilm))
	http.HandleFunc("/films/GetFilmByFragement", film.GetFilmByFragmentHandler(&newFilm))
	http.HandleFunc("/films/DeleteFilm", film.DeleteFilmHandler(&newFilm))
	http.HandleFunc("/films/GetFilmsByRate", film.SortFilmsByRateHandler(&newFilm))
	http.HandleFunc("/films/GetFilmsByName", film.SortFilmsByNameHandler(&newFilm))
	http.HandleFunc("/films/GetFilmsByDate", film.SortFilmsByDateHandler(&newFilm))

	//Server
	fmt.Println("Server is working  at http://127.0.0.1:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Start server error")
	}

}
