package film

import "sort"

type ByRate []Film

func (a ByRate) Len() int           { return len(a) }
func (a ByRate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRate) Less(i, j int) bool { return a[i].Rate > a[j].Rate }

func SortFilmsByRate(films []Film) []Film {
	sort.Sort(ByRate(films))
	return films
}

type ByName []Film

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func SortFilmsByName(films []Film) []Film {
	sort.Sort(ByName(films))
	return films
}

type ByDate []Film

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Enterdate.After(a[j].Enterdate) }

func SortFilmsByDate(films []Film) []Film {
	sort.Sort(ByDate(films))
	return films
}
