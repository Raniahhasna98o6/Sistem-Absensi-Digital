package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/absensi_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal connect database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database tidak merespon:", err)
	}

	DB = db
	log.Println("Database connected")
}
