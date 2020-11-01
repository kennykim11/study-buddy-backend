package main

import (
	"fmt"
	"time"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	google_id int unsigned,
	email varchar(20),
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
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
	db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/studybuddie")
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()
	tx.MustExec(schema)
	tx.MustExec("INSERT INTO users (given_name, family_name, email) VALUES (?, ?, ?)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.Commit()

	userSearchResult := User{}
	err = db.Get(&userSearchResult, "SELECT * FROM person WHERE given_name=?", "Jason")
	fmt.Printf("%#v\n", userSearchResult)
}
