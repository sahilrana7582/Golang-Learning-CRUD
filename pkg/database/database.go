package database

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
	"os"
)

func InitDB() (*sql.DB, error) {

	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("Error loading .env file")
	}

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}


func CreateTables(db *sql.DB) error {
    userTable := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(50) NOT NULL UNIQUE,
            email VARCHAR(100) NOT NULL UNIQUE,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )`
    
    if _, err := db.Exec(userTable); err != nil {
        return fmt.Errorf("error creating users table: %v", err)
    }

    // Create profiles table
    profileTable := `
        CREATE TABLE IF NOT EXISTS profiles (
            id SERIAL PRIMARY KEY,
            user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
            first_name VARCHAR(50),
            last_name VARCHAR(50),
            phone_number VARCHAR(20),
            address TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )`
    
    if _, err := db.Exec(profileTable); err != nil {
        return fmt.Errorf("error creating profiles table: %v", err)
    }

    return nil
}