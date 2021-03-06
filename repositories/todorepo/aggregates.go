// Code generated by nero, DO NOT EDIT.
package todorepo

import (
	"github.com/sf9v/nero/aggregate"
)

// AggFunc is an aggregate function
type AggFunc func(*aggregate.Aggregates)

// Avg is a average aggregate function
func Avg(col Column) AggFunc {
	return func(a *aggregate.Aggregates) {
		a.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Avg,
		})
	}
}

// Count is a count aggregate function
func Count(col Column) AggFunc {
	return func(a *aggregate.Aggregates) {
		a.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Count,
		})
	}
}

// Max is a max aggregate function
func Max(col Column) AggFunc {
	return func(a *aggregate.Aggregates) {
		a.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Max,
		})
	}
}

// Min is a min aggregate function
func Min(col Column) AggFunc {
	return func(a *aggregate.Aggregates) {
		a.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Min,
		})
	}
}

// Sum is a sum aggregate function
func Sum(col Column) AggFunc {
	return func(a *aggregate.Aggregates) {
		a.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.Sum,
		})
	}
}

// None is a none aggregate function
func None(col Column) AggFunc {
	return func(a *aggregate.Aggregates) {
		a.Add(&aggregate.Aggregate{
			Col: col.String(),
			Fn:  aggregate.None,
		})
	}
}
