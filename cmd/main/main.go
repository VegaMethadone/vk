package main

import (
	"fmt"
	"net/http"

	"vk/internal/authentication"
)

func main() {

	//used structers
	user := authentication.User{}

	//main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет от сервера!")
	})

	// user & authentication
	http.HandleFunc("/UserRegistration", authentication.RegistrationHandler(&user))
	http.HandleFunc("/EnterUser", authentication.EnterUserHandler(&user))
	http.HandleFunc("/ChangeUserData", authentication.ChangeUserDataHandler(&user))

	// films & acters

	fmt.Println("Server is working  at http://127.0.0.1:8080")
	//http://localhost:8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Start server error")
	}

}
