package db

import (
	"database/sql"

	_ "modernc.org/sqlite" // _ : to make sure that this package will not be removed when we save this file
)

// Start w/ uppercase to make this DB Global var
var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db") // open the db and assign to package var

	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10) // configure limit for number of connections to db
	DB.SetMaxIdleConns(5)  // configure limit for number of connections to db, when no one's using these connections at the moment

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		desription TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}

}
