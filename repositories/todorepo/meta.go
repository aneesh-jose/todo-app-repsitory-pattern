// Code generated by nero, DO NOT EDIT.
package todorepo

// Collection is the name of the collection
const Collection = "todos"

// Column is a Todo column
type Column int

// String returns the string representation of the Column
func (c Column) String() string {
	return [...]string{
		"id",
		"name",
		"description",
		"status",
		"user",
	}[c]
}

const (
	ColumnId Column = iota
	ColumnName
	ColumnDescription
	ColumnStatus
	ColumnUser
)
