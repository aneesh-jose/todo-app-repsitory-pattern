package main

import (
	"log"
	"os"

	"github.com/aneesh-jose/sample-todo/models"
	"github.com/sf9v/nero/gen"
)

func main() {
	files, err := gen.Generate((models.Todo{}).TodoSchema())
	if err != nil {
		log.Fatal(err)
	}

	basePath := "repositories/todorepo"
	err = os.Mkdir(basePath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err = file.Render(basePath)
		if err != nil {
			log.Fatal(err)
		}
	}

	files, err = gen.Generate((models.User{}).UserSchema())
	if err != nil {
		log.Fatal(err)
	}

	basePath = "repositories/userrepo"
	err = os.Mkdir(basePath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err = file.Render(basePath)
		if err != nil {
			log.Fatal(err)
		}
	}
}
