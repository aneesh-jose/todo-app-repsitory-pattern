package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aneesh-jose/sample-todo/repositories/todorepo"
	"github.com/aneesh-jose/sample-todo/repositories/userrepo"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	host := viper.Get("HOST")
	user := viper.Get("USER")
	password := viper.Get("PASSWORD")
	dbname := viper.Get("DBNAME")
	portStr, _ := viper.Get("PORT").(string)
	port, _ := strconv.Atoi(portStr)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	err = createUsersTable(db)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = dropUsersTable(db)
		if err != nil {
			log.Fatal(err)
		}
	}()

	userRepo := userrepo.NewPostgresRepository(db).Debug()
	ctx := context.Background()
	// create user 1
	user1Name, err := userRepo.Create(ctx, userrepo.NewCreator().Name("User 1").Password("samplepasscode").Username("user1"))
	if err != nil {
		log.Fatal(err)
	}

	// create user 2
	_, err = userRepo.Create(ctx, userrepo.NewCreator().Name("User 2").Password("samplepass").Username("user2"))
	if err != nil {
		log.Fatal(err)
	}

	// query user 1
	user1, err := userRepo.QueryOne(ctx, userrepo.NewQueryer().
		Where(userrepo.UsernameEq(user1Name)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", user1)

	// todo

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	err = createTodosTable(db)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = dropTodosTable(db)
		if err != nil {
			log.Fatal(err)
		}
	}()

	todoRepo := todorepo.NewPostgresRepository(db).Debug()
	ctx = context.Background()
	todoInput1, err := todoRepo.Create(
		ctx,
		todorepo.NewCreator().Id(
			int(time.Now().Unix())).
			Name("Todo 1").
			Description("Todo 1 description").
			User("user1").
			Status(false))
	if err != nil {
		log.Fatalf("Error1::::::%v", err.Error())
	}

	// create user 2
	time.Sleep(time.Second)
	_, err = todoRepo.Create(ctx,
		todorepo.NewCreator().
			Id(int(time.Now().Unix())).
			Name("Todo 2").
			Description("Todo 2 description").
			User("user2").
			Status(true))
	if err != nil {
		log.Fatalf("Error2::::::%v", err)
	}

	// query user 1
	todo1, err := todoRepo.QueryOne(ctx, todorepo.NewQueryer().
		Where(todorepo.IdEq(todoInput1)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", todo1)

	// update user 1

	_, err = todoRepo.Update(ctx, todorepo.NewUpdater().
		Status(true).
		Where(todorepo.IdEq(todoInput1)))
	if err != nil {
		log.Fatal(err)
	}

	// query user 1 again
	todo1, err = todoRepo.QueryOne(ctx, todorepo.
		NewQueryer().Where(todorepo.IdEq(todoInput1)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", todo1)

	// delete user 1
	_, err = todoRepo.Delete(ctx, todorepo.NewDeleter().
		Where(todorepo.IdEq(todoInput1)))
	if err != nil {
		log.Fatal(err)
	}

	// query remaining users
	users, err := todoRepo.Query(ctx, todorepo.NewQueryer().Limit(10))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", users[0])

}

func createTodosTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE todos (
		id int PRIMARY KEY,
		"name" VARCHAR(255) NOT NULL,
		"description" VARCHAR(255) NOT NULL,
		status bool NOT NULL,
		"user" varchar(255) NOT NULL
	)`)
	return err
}

func dropTodosTable(db *sql.DB) error {
	_, err := db.Exec(`drop table todos`)
	return err
}

func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE users (
		"username" varchar(255) PRIMARY KEY,
		"name" VARCHAR(255) NOT NULL,
		"password" VARCHAR(255) NOT NULL
	)`)
	return err
}

func dropUsersTable(db *sql.DB) error {
	_, err := db.Exec(`drop table users`)
	return err
}
