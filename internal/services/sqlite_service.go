package services

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	Db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := Db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	log.Println("Connected to SQLite database")

	if err := NewUserService(Db); err != nil {
		log.Printf("Error creating users table: %v", err)
	}
	if err := NewLocationService(Db).CreateTable(); err != nil {
		log.Printf("Error creating locations table: %v", err)
	}
	if err := NewSalonService(Db).CreateTable(); err != nil {
		log.Printf("Error creating salons table: %v", err)
	}
	if err := NewTableService(Db).CreateTable(); err != nil {
		log.Printf("Error creating tables table: %v", err)
	}
	if err := NewProductService(Db).CreateTable(); err != nil {
		log.Printf("Error creating products table: %v", err)
	}
}

func CloseDB() {
	if Db != nil {
		Db.Close()
		log.Println("Closed SQLite database connection")
	}
}
