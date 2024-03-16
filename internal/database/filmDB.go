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

	score := 1
	rate := 10.0 / float64(score)

	res, err := db.Exec(`
		INSERT INTO films (name, description, enterdate, rate, score) VALUES ($1, $2, $3, $4, $5)`,
		name, description, enterdate, rate, score)
	if err != nil {
		return false, err
	}

	_, data, err := FindFilmDB(name)
	if err != nil {
		return false, err
	}

	var id int
	var nameTmp string
	var descriptionTmp string
	var enterdateTmp time.Time
	var rateTmp float64
	var scoreTmp int

	for data.Next() {
		data.Scan(&id, &nameTmp, &descriptionTmp, &enterdateTmp, &rateTmp, &scoreTmp)
		for _, value := range acters {
			_, err := db.Exec(`
				INSERT INTO film_acters (film_id, acter_id) VALUES ($1, $2)`,
				id, value)
			if err != nil {
				return false, err
			}
		}
	}

	fmt.Println(res)
	return true, nil
}

func FindFilmDB(name string) (bool, *sql.Rows, error) {
	db, err := DBconnection()
	if err != nil {
		return false, nil, err
	}
	defer db.Close()

	res, err := db.Query(`
		SELECT id, name, description,  enterdate, rate, score
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

func ChangeFilmNameDB(id int, newName string) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`
		UPDATE films 
		SET name = $1
		WHERE id = $2`,
		newName, id)
	if err != nil {
		return false, err
	}

	fmt.Println(res)
	return true, nil
}

func ChangeFilmDescriptionNameDB(id int, newDescription string) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`
		UPDATE films
		SET description = $1
		WHERE id = $2`,
		newDescription, id)
	if err != nil {
		return false, err
	}

	fmt.Println(res)
	return true, nil
}

func ChangeFilmDateDB(id int, enterdate time.Time) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`
		UPDATE films
		SET enterdate = $1
		WHERE id = $2`,
		enterdate, id)
	if err != nil {
		return false, err
	}

	fmt.Println(res)
	return true, nil
}

func ChangeFilmRateDB(name string, score int) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	_, res, err := FindFilmDB(name)
	if err != nil {
		log.Fatal(err)
	}

	var id int
	var nameTmp string
	var descriptionTmp string
	var enterdateTmp time.Time
	var rateTmp float64
	var scoreTmp int

	for res.Next() {
		res.Scan(&id, &nameTmp, &descriptionTmp, &enterdateTmp, &rateTmp, &scoreTmp)

		rows, err := db.Exec(`
			UPDTAE films
			SET rate = $1,
			score = $2
			WHERE id = $3`,
			(scoreTmp+score)/10, scoreTmp+score, id)
		if err != nil {
			return false, err
		}
		fmt.Println(rows)
	}
	return true, nil
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

func ChangeFilmInfoDB(name, newName, newDescription string, newEnterdate time.Time, newScore int, changeName, changeDescription, changeEnterdate, changeScore bool) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	_, film, err := FindFilmDB(name)
	if err != nil {
		log.Fatalf("Error occured  %v", err)
	}

	var id int
	var nameTmp string
	var descriptionTmp string
	var enterdateTmp time.Time
	var rateTmp float64
	var scoreTmp int

	for film.Next() {
		film.Scan(&id, &nameTmp, &descriptionTmp, &enterdateTmp, &rateTmp, &scoreTmp)
	}

	if changeName {
		nameTmp = newName
	}
	if changeDescription {
		descriptionTmp = newDescription
	}
	if changeEnterdate {
		enterdateTmp = newEnterdate
	}
	if changeScore {
		scoreTmp += newScore
		rateTmp = float64(scoreTmp) / 10.0
	}

	res, err := db.Exec(`
		UPDATE films 
		SET name = $1,
		description = $2,
		enterdate = $3,
		rate = &4,
		score = $5`,
		nameTmp, descriptionTmp, enterdateTmp, rateTmp, scoreTmp)
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
