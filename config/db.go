package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, *sql.DB) {
	// Get DATABASE_URL from Heroku env
	dbURL := os.Getenv("DATABASE_URL")

	sqlDbExst, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDbExst,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to Ping to DB:", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatal("DB not reachable:", err)
	}

	fmt.Println("Database connected!")
	return db, sqlDB
}
