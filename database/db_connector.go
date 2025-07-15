package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"os"
	"pokerdegen/models"
)

func ConnectDB() (*sql.DB, error) {
	connStr := func() string {
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		dbname := os.Getenv("DB_NAME")
		sslmode := os.Getenv("DB_SSLMODE")

		if sslmode == "" {
			sslmode = "require"
		}

		return fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			user, password, host, port, dbname, sslmode,
		)
	}

	db, err := sql.Open("postgres", connStr())
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping error: %w", err)
	}

	return db, nil
}

func InsertUser(db *sql.DB, username string, password string, diamonds int) error {
	query := `
		INSERT INTO users (username, password, diamonds)
		VALUES ($1, $2, $3)
	`
	_, err := db.Exec(query, username, password, diamonds)
	if err != nil {
		return fmt.Errorf("InsertUser error: %w", err)
	}

	return nil
}

func FetchUser(db *sql.DB, username string, password string) (*models.User, error) {
	query := `
		SELECT id, username, password, diamonds
		FROM users
		WHERE username = $1 and password = $2
	`

	row := db.QueryRow(query, username, password)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Diamonds)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("query error: %w", err)
	}

	return &user, nil
}