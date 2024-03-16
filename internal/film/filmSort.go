package film

import "sort"

type ByRate []film

func (a ByRate) Len() int           { return len(a) }
func (a ByRate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRate) Less(i, j int) bool { return a[i].rate > a[j].rate }

func SortFilmsByRate(films []film) []film {
	sort.Sort(ByRate(films))
	return films
}

type ByName []film

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].name < a[j].name }

func SortFilmsByName(films []film) []film {
	sort.Sort(ByName(films))
	return films
}

type ByDate []film

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].enterdate.After(a[j].enterdate) }

func SortFilmsByDate(films []film) []film {
	sort.Sort(ByDate(films))
	return films
}
