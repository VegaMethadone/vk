package film

import (
	"time"
)

type Film struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Enterdate   time.Time `json:"enterdate"`
	Rate        float64   `json:"rate"`
	Score       int       `json:"score"`
	Votes       int       `json:"votes"`
	Acters      []int     `json:"acters"`
}
