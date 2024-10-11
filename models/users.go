package models

import "database/sql"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u *User) GetUser() (*User, error) {

	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "messages.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Prepare the SQL statement
	stmt, err := db.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	row := stmt.QueryRow(&u.ID)
	err = row.Scan(&u.ID, &u.Name)
	if err != nil {
		return nil, err
	}

	return u, nil

}

func (u *User) GetAllUsers() ([]User, error) {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "messages.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	// Prepare the SQL statement
	stmt, err := db.Prepare("SELECT * FROM users")
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

	// Create a slice of User structs to hold the results
	users := []User{}

	// Iterate over the rows returned by the query
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *User) CreateUser() error {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "messages.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Check if the user already exists
	_, err = u.GetUser()
	if err == nil {
		return nil
	}

	// Prepare the SQL statement
	stmt, err := db.Prepare("INSERT INTO users (name) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(u.Name)

	if err != nil {
		return err
	}

	return nil
}
