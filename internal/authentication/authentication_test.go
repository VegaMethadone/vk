package authentication

import (
	"log"
	"testing"
)

func TestFindUser(t *testing.T) {
	expectedUser := &User{
		Login:    "Pog",
		Password: "123123",
		Access:   1,
	}

	newUser := &User{}

	result, newUser := newUser.FindUser(expectedUser.Login, expectedUser.Password)
	if !result {
		log.Fatalf("Did not find user %v\n", result)
	}

	log.Fatalf("Expected %v, %v, GOT: %v, %v\n", expectedUser.Login, expectedUser.Password, newUser.Login, newUser.Password)
}
