package models

import "github.com/sf9v/nero"

type Todo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	User        string `json:"username"`
}

func (t Todo) TodoSchema() *nero.Schema {
	return (nero.NewSchemaBuilder(&t)).
		PkgName("todorepo").
		Collection("todos").
		Identity(
			nero.NewColumnBuilder("id", t.Id).
				StructField("Id").Build(),
		).
		Columns(
			nero.NewColumnBuilder("name", t.Name).Build(),
			nero.NewColumnBuilder("description", t.Description).Build(),
			nero.NewColumnBuilder("status", t.Status).Build(),
			nero.NewColumnBuilder("user", t.User).Build(),
		).Build()
}
