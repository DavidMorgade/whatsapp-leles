package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "messages.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the users table
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE
	);`
	_, err = db.Exec(createUsersTable)
	if err != nil {
		log.Fatal(err)
	}

	// Create the messages table
	createMessagesTable := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		message TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`
	_, err = db.Exec(createMessagesTable)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database and tables created successfully")
}
