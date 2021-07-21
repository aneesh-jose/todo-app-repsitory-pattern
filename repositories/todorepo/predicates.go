// Code generated by nero, DO NOT EDIT.
package todorepo

import (
	"github.com/sf9v/nero/comparison"
)

// PredFunc is a predicate function
type PredFunc func(*comparison.Predicates)

// IdEq applies "equal" operator on "id" column
func IdEq(id int) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Eq,
			Arg: id,
		})
	}
}

// IdNotEq applies "not equal" operator on "id" column
func IdNotEq(id int) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotEq,
			Arg: id,
		})
	}
}

// IdGt applies "greater than" operator on "id" column
func IdGt(id int) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Gt,
			Arg: id,
		})
	}
}

// IdGtOrEq applies "greater than or equal" operator on "id" column
func IdGtOrEq(id int) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.GtOrEq,
			Arg: id,
		})
	}
}

// IdLt applies "less than" operator on "id" column
func IdLt(id int) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.Lt,
			Arg: id,
		})
	}
}

// IdLtOrEq applies "less than or equal" operator on "id" column
func IdLtOrEq(id int) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.LtOrEq,
			Arg: id,
		})
	}
}

// IdIn applies "in" operator on "id" column
func IdIn(ids ...int) PredFunc {
	args := []interface{}{}
	for _, v := range ids {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.In,
			Arg: args,
		})
	}
}

// IdNotIn applies "not in" operator on "id" column
func IdNotIn(ids ...int) PredFunc {
	args := []interface{}{}
	for _, v := range ids {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "id",
			Op:  comparison.NotIn,
			Arg: args,
		})
	}
}

// NameEq applies "equal" operator on "name" column
func NameEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Eq,
			Arg: name,
		})
	}
}

// NameNotEq applies "not equal" operator on "name" column
func NameNotEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.NotEq,
			Arg: name,
		})
	}
}

// NameGt applies "greater than" operator on "name" column
func NameGt(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Gt,
			Arg: name,
		})
	}
}

// NameGtOrEq applies "greater than or equal" operator on "name" column
func NameGtOrEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.GtOrEq,
			Arg: name,
		})
	}
}

// NameLt applies "less than" operator on "name" column
func NameLt(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.Lt,
			Arg: name,
		})
	}
}

// NameLtOrEq applies "less than or equal" operator on "name" column
func NameLtOrEq(name string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.LtOrEq,
			Arg: name,
		})
	}
}

// NameIn applies "in" operator on "name" column
func NameIn(names ...string) PredFunc {
	args := []interface{}{}
	for _, v := range names {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.In,
			Arg: args,
		})
	}
}

// NameNotIn applies "not in" operator on "name" column
func NameNotIn(names ...string) PredFunc {
	args := []interface{}{}
	for _, v := range names {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "name",
			Op:  comparison.NotIn,
			Arg: args,
		})
	}
}

// DescriptionEq applies "equal" operator on "description" column
func DescriptionEq(description string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "description",
			Op:  comparison.Eq,
			Arg: description,
		})
	}
}

// DescriptionNotEq applies "not equal" operator on "description" column
func DescriptionNotEq(description string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "description",
			Op:  comparison.NotEq,
			Arg: description,
		})
	}
}

// DescriptionGt applies "greater than" operator on "description" column
func DescriptionGt(description string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "description",
			Op:  comparison.Gt,
			Arg: description,
		})
	}
}

// DescriptionGtOrEq applies "greater than or equal" operator on "description" column
func DescriptionGtOrEq(description string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "description",
			Op:  comparison.GtOrEq,
			Arg: description,
		})
	}
}

// DescriptionLt applies "less than" operator on "description" column
func DescriptionLt(description string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "description",
			Op:  comparison.Lt,
			Arg: description,
		})
	}
}

// DescriptionLtOrEq applies "less than or equal" operator on "description" column
func DescriptionLtOrEq(description string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "description",
			Op:  comparison.LtOrEq,
			Arg: description,
		})
	}
}

// DescriptionIn applies "in" operator on "description" column
func DescriptionIn(descriptions ...string) PredFunc {
	args := []interface{}{}
	for _, v := range descriptions {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "description",
			Op:  comparison.In,
			Arg: args,
		})
	}
}

// DescriptionNotIn applies "not in" operator on "description" column
func DescriptionNotIn(descriptions ...string) PredFunc {
	args := []interface{}{}
	for _, v := range descriptions {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "description",
			Op:  comparison.NotIn,
			Arg: args,
		})
	}
}

// StatusEq applies "equal" operator on "status" column
func StatusEq(status bool) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "status",
			Op:  comparison.Eq,
			Arg: status,
		})
	}
}

// StatusNotEq applies "not equal" operator on "status" column
func StatusNotEq(status bool) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "status",
			Op:  comparison.NotEq,
			Arg: status,
		})
	}
}

// StatusGt applies "greater than" operator on "status" column
func StatusGt(status bool) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "status",
			Op:  comparison.Gt,
			Arg: status,
		})
	}
}

// StatusGtOrEq applies "greater than or equal" operator on "status" column
func StatusGtOrEq(status bool) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "status",
			Op:  comparison.GtOrEq,
			Arg: status,
		})
	}
}

// StatusLt applies "less than" operator on "status" column
func StatusLt(status bool) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "status",
			Op:  comparison.Lt,
			Arg: status,
		})
	}
}

// StatusLtOrEq applies "less than or equal" operator on "status" column
func StatusLtOrEq(status bool) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "status",
			Op:  comparison.LtOrEq,
			Arg: status,
		})
	}
}

// StatusIn applies "in" operator on "status" column
func StatusIn(statuses ...bool) PredFunc {
	args := []interface{}{}
	for _, v := range statuses {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "status",
			Op:  comparison.In,
			Arg: args,
		})
	}
}

// StatusNotIn applies "not in" operator on "status" column
func StatusNotIn(statuses ...bool) PredFunc {
	args := []interface{}{}
	for _, v := range statuses {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "status",
			Op:  comparison.NotIn,
			Arg: args,
		})
	}
}

// UserEq applies "equal" operator on "user" column
func UserEq(user string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "user",
			Op:  comparison.Eq,
			Arg: user,
		})
	}
}

// UserNotEq applies "not equal" operator on "user" column
func UserNotEq(user string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "user",
			Op:  comparison.NotEq,
			Arg: user,
		})
	}
}

// UserGt applies "greater than" operator on "user" column
func UserGt(user string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "user",
			Op:  comparison.Gt,
			Arg: user,
		})
	}
}

// UserGtOrEq applies "greater than or equal" operator on "user" column
func UserGtOrEq(user string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "user",
			Op:  comparison.GtOrEq,
			Arg: user,
		})
	}
}

// UserLt applies "less than" operator on "user" column
func UserLt(user string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "user",
			Op:  comparison.Lt,
			Arg: user,
		})
	}
}

// UserLtOrEq applies "less than or equal" operator on "user" column
func UserLtOrEq(user string) PredFunc {
	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "user",
			Op:  comparison.LtOrEq,
			Arg: user,
		})
	}
}

// UserIn applies "in" operator on "user" column
func UserIn(users ...string) PredFunc {
	args := []interface{}{}
	for _, v := range users {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "user",
			Op:  comparison.In,
			Arg: args,
		})
	}
}

// UserNotIn applies "not in" operator on "user" column
func UserNotIn(users ...string) PredFunc {
	args := []interface{}{}
	for _, v := range users {
		args = append(args, v)
	}

	return func(pb *comparison.Predicates) {
		pb.Add(&comparison.Predicate{
			Col: "user",
			Op:  comparison.NotIn,
			Arg: args,
		})
	}
}