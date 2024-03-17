package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func AddNewFilmDB(name, description string, enterdate time.Time, acters []int) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	rate := 0.0
	score := 0
	votes := 0

	res, err := db.Exec(`
		INSERT INTO films (name, description, enterdate, rate, score, votes) VALUES ($1, $2, $3, $4, $5, $6)`,
		name, description, enterdate, rate, score, votes)
	if err != nil {
		fmt.Println("Error duing inserting")
		return false, err
	}

	_, film, err := FindFilmDB(name)
	if err != nil {
		fmt.Println("Error during finding film")
		return false, err
	}
	var filmID int

	for film.Next() {
		film.Scan(&filmID, &name, &description, &enterdate, &rate, &score, &votes)
	}

	//PRINT FILM ID
	fmt.Println("Film ID:", filmID)

	for _, acterID := range acters {
		_, err = FindActerByIdDB(acterID)
		if err != nil {
			continue
		} else {
			//COMMENT
			fmt.Println("Found acter id,", acterID)
			//COMMENT
			res, err := db.Exec(`
				INSERT INTO film_acters (film_id, acter_id) VALUES ($1, $2)`, filmID, acterID)
			if err != nil {
				return false, err
			}
			fmt.Println(res)
		}
	}
	fmt.Println(res)
	return true, nil
}

//

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

func FindFilmByIdDB(id int) (bool, *sql.Rows, error) {
	db, err := DBconnection()
	if err != nil {
		return false, nil, err
	}
	defer db.Close()

	res, err := db.Query(`
		SELECT id, name, description,  enterdate, rate, score, votes
		FROM filmse
		WHERE id = $1`, id)
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

func ChangeFilmInfoDB(id int, name, description string, enterdate time.Time, score, votes int, acters []int) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	result, err := db.Exec(`
		UPDATE films
		SET name= $1, 
		description= $2, 
		enterdate= $3, 
		rate= $4, 
		score=  $5, 
		votes= $6 
		WHERE id = $7`,
		name, description, enterdate, score/votes, score, votes, id)
	if err != nil {
		fmt.Println("Error during setting new film data")
		return false, err
	}
	fmt.Println(result)

	DeleteFilmActers(id)

	for _, acterId := range acters {
		_, err := db.Exec(`
			INSERT INTO film_acters (film_id, acter_id) VALUES ($1, $2)`, id, acterId)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func DeleteFilmActers(filmID int) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`
		DELETE FROM film_acters WHERE film_id =  $1`, filmID)
	if err != nil {
		return false, err
	}
	fmt.Println(res)
	return true, nil
}
