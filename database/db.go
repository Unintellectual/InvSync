package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	if user == "" || password == "" || database == "" {
		log.Fatal("Database credentials are missing in environment variables")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}
	fmt.Println("Connected to MySQL")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Database connection closed")
	}
}
