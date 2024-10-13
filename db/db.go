package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
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
		id STRING PRIMARY KEY NOT NULL,
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

func CreateWaDB() (deviceStore *store.Device, clientLog waLog.Logger, err error) {

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	dsn := "file:whatsapp.db?_foreign_keys=on"
	container, err := sqlstore.New("sqlite3", dsn, dbLog)
	if err != nil {
		return nil, nil, err
	}

	deviceStore, err = container.GetFirstDevice()
	if err != nil {
		return nil, nil, err
	}

	clientLog = waLog.Stdout("Client", "INFO", true)

	return deviceStore, clientLog, nil
}
