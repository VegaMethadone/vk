package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func AddNewActerDB(name, sex string, dateOfBirth time.Time) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(
		`INSERT INTO acters (name,  sex,  dateofbirth) VALUES ($1, $2, $3)`,
		name, sex, dateOfBirth)

	if err != nil {
		return false, err
	}
	log.Println("Added new acter db", res)
	return true, nil

}

func FindActerDB(name string) (bool, *sql.Rows, error) {
	db, err := DBconnection()
	if err != nil {
		return false, nil, err
	}
	defer db.Close()

	res, err := db.Query(`
		SELECT id, name,  sex, dateOfBirth
		FROM acters
		WHERE name =  $1`,
		name)
	if err != nil {
		return false, nil, err
	}
	log.Println("Found acter db:", res)
	return true, res, nil

}

func ChangeActerNameDB(id int, name string) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`
		UPDATE acters
        SET name = $1
        WHERE id = $2`,
		name, id)
	if err != nil {
		return false, err
	}
	log.Println("Changed acter name db:", res)
	return true, nil
}

func ChangeActerSexDB(id int, sex string) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`
		UPDATE acters
        SET sex = $1
        WHERE id = $2`,
		sex, id)
	if err != nil {
		return false, err
	}
	fmt.Println(res)
	return true, nil
}

func ChangeActerDateOfBirthDB(id int, date time.Time) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`
		UPDATE acters
        SET dateOfBirth = $1
        WHERE id = $2`,
		date, id)
	if err != nil {
		return false, err
	}
	log.Println("Changed acter date DB", res)
	return true, nil
}

func ChangeActerAllDB(id int, name, sex string, date time.Time) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`
		UPDATE  acters
		SET name = $1,
		sex = $2,
		dateOfBirth = $3
		WHERE id = $4`,
		name, sex, date, id)
	if err != nil {
		return false, err
	}
	log.Println("Changed acter info DB", res)
	return true, nil
}

func DeleteActerInfoDB(id int) (bool, error) {
	db, err := DBconnection()
	if err != nil {
		return false, err
	}
	defer db.Close()

	res, err := db.Exec(`
		DELETE FROM acters
		WHERE id  =  $1`,
		id)
	if err != nil {
		return false, err
	}
	log.Println("Deleted  acter info DB:", res)
	return true, nil
}

func GetAllActersDB() (*sql.Rows, error) {
	db, err := DBconnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT * FROM acters`)
	if err != nil {
		return nil, err
	}

	return rows, nil

}

func GetALLActerFilmsDB(id int) (*sql.Rows, error) {
	db, err := DBconnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT film_id FROM film_acters WHERE acter_id  = $1`, id)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func FindActerByIdDB(id int) (*sql.Rows, error) {
	db, err := DBconnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	res, err := db.Query(`
		SELECT id, name,  sex, dateOfBirth
		FROM acters
		WHERE id =  $1`,
		id)
	if err != nil {
		return nil, err
	}
	log.Println("Found acter  db by ID:", res)
	return res, nil

}

func DeleteAllActerFilmsDB(id int) error {
	db, err := DBconnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`
		DELETE FROM film_acters WHERE acter_id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
