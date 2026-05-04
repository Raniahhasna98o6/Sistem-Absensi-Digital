package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// Ganti username dan password ini dengan kredensial Azure MySQL lu
	// DSN format: username:password@tcp(host:port)/dbname
	dsn := "sean:Bean2080!!!!@tcp(db-absensi-telyu.mysql.database.azure.com:3306)/db_absensi?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal konfigurasi database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database tidak merespon, pastikan host dan kredensial Azure benar:", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	log.Println("Berhasil nyambung ke Database Azure MySQL!")
	DB = db
}
