package database

import (
	"database/sql"
	"fmt"
	"log"
)

func FindFilmDB(name string) (bool, *sql.Rows, error) {
	db, err := DBconnection()
	if err != nil {
		return false, nil, err
	}
	defer db.Close()

	res, err := db.Query(`
		SELECT id, name, description,  enterdate, rate, score, votes
		FROM films
		WHERE name =  $1`,
		name)
	if err != nil {
		return false, nil, err
	}
	//fmt.Println(res)
	return true, res, nil

}

func FindFilmActersDB(id int) (bool, *sql.Rows, error) {
	db, err := DBconnection()
	if err != nil {
		return false, nil, err
	}
	defer db.Close()

	res, err := db.Query(`
		SELECT acter_id FROM film_acters
		WHERE film_id = $1`,
		id)
	if err != nil {
		return false, nil, err
	}
	//fmt.Println(res)
	return true, res, nil
}

func DeleteFilmDB(id int) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`
		DELETE FROM films WHERE id = $1`, id)
	if err != nil {
		return false, err
	}
	fmt.Println(res)
	return true, nil
}

///

func FindFilmByFragment(fragment string) (bool, *sql.Rows, error) {
	db, err := DBconnection()
	if err != nil {
		return false, nil, err
	}
	defer db.Close()

	res, err := db.Query(`
		SELECT id, name, description, enterdate, rate, score FROM films
		WHERE name LIKE '%' || $1 || '%';`, fragment)
	if err != nil {
		log.Fatalf("Error occured %v", err)
	}

	fmt.Println(res)
	return true, res, nil

}

func GetAllFilmsDB() (*sql.Rows, error) {
	db, err := DBconnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	res, err := db.Query(`
		SELECT * FROM films`)
	if err != nil {
		log.Fatalf("Unexpected  error during getting  films %v", err)
		return nil, err
	}
	return res, nil
}

func GetAllFilmActersDB(id int) (*sql.Rows, error) {
	db, err := DBconnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT acter_id FROM film_acters WHERE film_id  = $1`, id)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
