// Code generated by nero, DO NOT EDIT.
package userrepo

// Collection is the name of the collection
const Collection = "users"

// Column is a User column
type Column int

// String returns the string representation of the Column
func (c Column) String() string {
	return [...]string{
		"username",
		"name",
		"password",
	}[c]
}

const (
	ColumnUsername Column = iota
	ColumnName
	ColumnPassword
)
