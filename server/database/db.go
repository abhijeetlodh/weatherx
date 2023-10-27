package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {

	dbURL := os.Getenv("DB")

	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		os.Exit(1)
	}
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func GetUserByID(userID int) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, email, firstname, password, created_at, updated_at FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Email, &user.FirstName, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, email, firstname, password, created_at, updated_at FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.FirstName, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func InsertUser(user *User) error {
	_, err := db.Exec("INSERT INTO users (email, firstname, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", user.Email, user.FirstName, user.Password, user.CreatedAt, user.UpdatedAt)
	return err
}

func UpdateUser(user *User) error {
	_, err := db.Exec("UPDATE users SET email = $1, firstname = $2, password = $3, updated_at = $4 WHERE id = $5", user.Email, user.FirstName, user.Password, user.UpdatedAt, user.ID)
	return err
}

func DeleteUser(userID int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	return err
}
