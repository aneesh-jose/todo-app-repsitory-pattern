// Code generated by nero, DO NOT EDIT.
package todorepo

import (
	"github.com/sf9v/nero/sort"
)

// SortFunc is a sort function
type SortFunc func(*sort.Sorts)

// Asc sorts in ascending order
func Asc(col Column) SortFunc {
	return func(s *sort.Sorts) {
		s.Add(&sort.Sort{
			Col:       col.String(),
			Direction: sort.Asc,
		})
	}
}

// Desc sorts in descending order
func Desc(col Column) SortFunc {
	return func(s *sort.Sorts) {
		s.Add(&sort.Sort{
			Col:       col.String(),
			Direction: sort.Desc,
		})
	}
}
