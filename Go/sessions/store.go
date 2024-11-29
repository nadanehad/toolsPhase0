package sessions

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var SessionStore = struct {
	sync.RWMutex
	Sessions map[string]uint
}{Sessions: make(map[string]uint)}

// InitDB initializes the database connection
func InitDB() *sql.DB {
	// Read environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST") // Use "database" as the host in Docker
	dbName := os.Getenv("DB_NAME")

	// Create the connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)

	// Open the database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping the database: %v", err))
	}

	fmt.Println("Connected to the database successfully!")
	return db
}
