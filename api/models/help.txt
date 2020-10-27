package main

import (
	"fmt"
	"time"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE users (
	google_id int unsigned,
	email varchar(20),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    given_name varchar(20),
    family_name varchar(20),
    bio_description varchar(200),
	KEY (google_id)
);`

type User struct {
	googleID       uint32
	email          string
	givenName      string
	familyName     string
	bioDescription string
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      time.Time
}

func main() {
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/random_testing")
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO users (given_name, family_name, email) VALUES (?, ?, ?)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	//tx.NamedExec("INSERT INTO users (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
	tx.Commit()

	// You can also get a single result, a la QueryRow
	userSearchResult := User{}
	err = db.Get(&userSearchResult, "SELECT * FROM person WHERE given_name=?", "Jason")
	fmt.Printf("%#v\n", userSearchResult)
}
