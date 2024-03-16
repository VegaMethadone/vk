package database

import (
	"database/sql"
	"fmt"
)

func NewUserDB(login, password string, access int) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`
		INSERT INTO users (login, password, access)
		VALUES ($1, $2, $3)`,
		login, password, access)
	if err != nil {
		return false, err
	}
	fmt.Println(res)
	return true, nil
}

func FindUserDB(login, password string) (bool, *sql.Rows, error) {
	db, err := DBconnection()
	if err != nil {
		return false, nil, err
	}
	defer db.Close()

	res, err := db.Query(`
		SELECT id, login,password, access FROM users
		WHERE login = $1 AND password = $2`,
		login, password)
	if err != nil {
		return false, nil, err
	}
	//fmt.Println(res)
	return true, res, nil
}

func ChangePasswordOrLoginDB(login, password string, access, id int) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`
		UPDATE users
		SET login = $1,
			password = $2,
			access = $3
		WHERE  id = $4`,
		login, password, access, id)
	if err != nil {
		return false, err
	}
	fmt.Println(res)
	return true, nil
}
