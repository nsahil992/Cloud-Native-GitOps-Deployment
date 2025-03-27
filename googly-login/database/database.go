package database

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

// InitDB initializes database connection
func InitDB() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to database")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

// UserExists checks if a user with the given email exists
func UserExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
	err := DB.QueryRow(query, email).Scan(&exists)
	return exists, err
}

// CreateUser inserts a new user into the database
func CreateUser(name, email, hashedPassword string) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err := DB.Exec(query, name, email, hashedPassword)
	return err
}

// GetUserByEmail retrieves a user by email for login verification
func GetUserByEmail(email string) (int, string, string, error) {
	var id int
	var name, password string
	query := "SELECT id, name, password FROM users WHERE email = $1"
	err := DB.QueryRow(query, email).Scan(&id, &name, &password)
	return id, name, password, err
}
