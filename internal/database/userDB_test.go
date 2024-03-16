package database

import (
	"fmt"
	"testing"
)

type testUser struct {
	id       int
	login    string
	password string
	access   int
}

func TestNewUserDB(t *testing.T) {
	userOne := &testUser{
		id:       1,
		login:    "Pog",
		password: "123123",
		access:   1,
	}
	userTwo := &testUser{
		id:       2,
		login:    "Kek",
		password: "321321",
		access:   0,
	}

	db, err := DBconnection()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	resOne, err := NewUserDB(userOne.login, userOne.password, userOne.access)
	if err != nil {
		t.Fatalf("unexpected error UserOne: %v", err)
	}
	if !resOne {
		t.Errorf("userOne was not added")
	}

	resTwo, err := NewUserDB(userTwo.login, userTwo.password, userTwo.access)
	if err != nil {
		t.Fatalf("unexpected error UserTwo: %v", err)
	}
	if !resTwo {
		t.Errorf("userTwo was not added")
	}

}

func TestFindUserDB(t *testing.T) {
	userOne := &testUser{
		id:       1,
		login:    "Pog",
		password: "123123",
		access:   1,
	}
	userTwo := &testUser{
		id:       2,
		login:    "Kek",
		password: "321321",
		access:   0,
	}
	arr := []testUser{*userOne, *userTwo}

	db, err := DBconnection()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	for _, value := range arr {
		flag, res, err := FindUserDB(value.login, value.password)
		if err != nil {
			t.Fatalf("unexpected error UserOne: %v", err)
		}
		if !flag {
			t.Errorf("expected user to be found, but was not")
		}

		defer res.Close()

		for res.Next() {
			var id int
			var login string
			var password string
			var access int
			_ = res.Scan(&id, &login, &password, &access)
			fmt.Println("EXPECTED:", value.login, " GOT:", login)
		}

	}

}

func TestChangePasswordOrLoginDB(t *testing.T) {
	userTwo := &testUser{
		id:       2,
		login:    "kek",
		password: "321321",
		access:   0,
	}
	db, err := DBconnection()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	resTwo, err := ChangePasswordOrLoginDB("rol", userTwo.password, userTwo.access, userTwo.id)
	if err != nil {
		t.Fatalf("unexpected error UserTwo: %v", err)
	}
	if !resTwo {
		t.Errorf("Did not change data in userTwo")
	}

	flag, res, err := FindUserDB("rol", userTwo.password)
	if err != nil {
		t.Fatalf("unexpected error UserOne: %v", err)
	}
	if !flag {
		t.Errorf("expected user to be found, but was not found")
	}

	defer res.Close()

	for res.Next() {
		var login string
		_ = res.Scan(&login)
		fmt.Println("EXPECTED:", userTwo.login, " GOT:", login)
	}

}
