package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	db, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect with database")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to database successful")

	DB = db

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	);`
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		log.Fatal(err)
		// panic("Couldn't create events table")
	}
}
