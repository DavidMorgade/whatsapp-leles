package models

import "database/sql"

type Message struct {
	ID        int    `json:"id"`
	UserID    string `json:"user_id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func (m *Message) GetMessage() (*Message, error) {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "messages.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Prepare the SQL statement
	stmt, err := db.Prepare("SELECT * FROM messages WHERE id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	row := stmt.QueryRow(&m.ID)
	err = row.Scan(&m.ID, &m.UserID, &m.Message, &m.Timestamp)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Message) GetAllMessages() ([]Message, error) {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "messages.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	// Prepare the SQL statement
	stmt, err := db.Prepare("SELECT * FROM messages")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice of Message structs to hold the results
	messages := []Message{}

	// Iterate over the rows
	for rows.Next() {
		var m Message
		err := rows.Scan(&m.ID, &m.UserID, &m.Message, &m.Timestamp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

func (m *Message) SaveMessage() error {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "messages.db")
	if err != nil {
		return err
	}

	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO messages(user_id, message) values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(&m.UserID, &m.Message)
	if err != nil {
		return err
	}

	return nil
}
