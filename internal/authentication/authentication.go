package authentication

import (
	"log"
	"strings"

	"vk/internal/database"
)

func findSubstring(str, target string) bool {
	str = strings.ToLower(str)

	index := strings.Index(str, target)
	if index == -1 {
		return false
	}
	return true
}

func (u *User) CreateUser(login, password string, access int) bool {
	if findSubstring(login, "drop") || findSubstring(password, "drop") {
		log.Println("found a forbidden string - Creaue User")
		return false
	}

	result, err := database.NewUserDB(login, password, access)
	if err != nil {
		log.Printf("Error occurred during registration %v", err)
		return false
	}
	return result
}

func (u *User) FindUser(login, password string) (bool, *User) {
	_, userData, err := database.FindUserDB(login, password)
	if err != nil {
		log.Printf("Error occurred during finding user %v", err)
		return false, nil
	}

	foundUser := new(User)
	for userData.Next() {
		userData.Scan(&foundUser.Id, &foundUser.Login, &foundUser.Password, &foundUser.Access)

		defer userData.Close()
	}

	if foundUser.Login == login && foundUser.Password == password {
		log.Println("Found user:", login)
		return true, foundUser
	}

	return false, nil
}

func (u *User) ChangeUserData(id, access int, login, password string) (bool, *User) {
	result, err := database.ChangePasswordOrLoginDB(login, password, access, id)
	if err != nil {
		log.Printf("Unexpected error uccurred during user changing %v", err)
		return false, nil
	}
	newUser := &User{
		Id:       id,
		Login:    login,
		Password: password,
		Access:   access,
	}
	log.Println(newUser)

	return result, newUser
}
