package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	    id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	    email text NOT NULL UNIQUE,
	    password text NOT NULL
	)
`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println("Could not create users table:", err)
		panic("Could not create users.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	    id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    dateTime DATETIME NOT NULL,
	    creator_id integer NOT NULL,
	    FOREIGN KEY(creator_id) REFERENCES users(id)                              
	)
`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println("Could not create events table:", err)
		panic("Could not create events table.")
	}
}
