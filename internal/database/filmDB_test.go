package database

import (
	"fmt"
	"testing"
	"time"
)

type testFilm struct {
	//id          int
	name        string
	description string
	enterdate   time.Time
	rate        float64
	score       int
	acters      []int
}

/*
func TestAddNewFilmDB(t *testing.T) {
	filmOne := &testFilm{
		name:        "Drive",
		description: "Великолепный водитель – при свете дня он выполняет каскадерские трюки на съёмочных площадках Голливуда, а по ночам ведет рискованную игру. Но один опасный контракт – и за его жизнь назначена награда. Теперь, чтобы остаться в живых и спасти свою очаровательную соседку, он должен делать то, что умеет лучше всего – виртуозно уходить от погони.",
		enterdate:   time.Date(2011, time.May, 20, 0, 0, 0, 0, time.UTC),
		rate:        7.3,
		score:       73,
		acters:      []int{1},
	}

	filmTwo := &testFilm{
		name:        "Pusher",
		description: "У пушера Франка крупные неприятности. Он должен 50 тысяч поставщикам-югославам, а вдобавок к этому во время продажи героина приезжим шведам налетела полиция, и весь взятый в долг товар пришлось выбросить в пруд. Друг Тонни, с которым он вместе работал, во всем признался полиции. Но товар выброшен, улик нет, и Франка выпустили. Теперь его долг поставщику Мило — 180 тысяч",
		enterdate:   time.Date(1996, time.August, 30, 0, 0, 0, 0, time.UTC),
		rate:        7.0,
		score:       70,
		acters:      []int{2},
	}

	films := []testFilm{*filmOne, *filmTwo}

	for key, value := range films {
		res, err := AddNewFilmDB(value.name, value.description, value.enterdate, value.acters)
		if err != nil {
			t.Fatalf("Unexpected error occured %v", err)
		}
		if !res {
			t.Errorf("Acter %d was not added", key)
		}
		fmt.Println(res)

	}
}
*/

func TestFindFilmByFragment(t *testing.T) {
	fragment := "Driv"

	_, res, err := FindFilmByFragmentDB(fragment)
	if err != nil {
		t.Fatalf("Did not  find %v", err)
	}

	var id int
	var nameTmp string
	var descriptionTmp string
	var enterdateTmp time.Time
	var rateTmp float64
	var scoreTmp int
	var votes int

	for res.Next() {
		res.Scan(&id, &nameTmp, &descriptionTmp, &enterdateTmp, &rateTmp, &scoreTmp, &votes)

		fmt.Println("EXPECTED:  Driver", "\nGOT ", nameTmp)
		fmt.Println(id, nameTmp, descriptionTmp, enterdateTmp, rateTmp, scoreTmp, votes)

	}

}
