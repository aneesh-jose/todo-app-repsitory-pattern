package models

import "github.com/sf9v/nero"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (u User) UserSchema() *nero.Schema {
	return (nero.NewSchemaBuilder(&u)).
		PkgName("userrepo").
		Collection("users").
		Identity(
			nero.NewColumnBuilder("username", u.Username).Build(),
		).
		Columns(
			nero.NewColumnBuilder("name", u.Name).Build(),
			nero.NewColumnBuilder("password", u.Password).Build(),
		).Build()
}
