package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// Ganti nama databasenya jadi db-absensi-telyu sesuai dengan yang di Azure
	dsn := "sean:Bean2080%21%21%21%21@tcp(db-absensi-telyu.mysql.database.azure.com:3306)/db-absensi-telyu?parseTime=true&tls=skip-verify"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Gagal konfigurasi database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("Database tidak merespon (Cek Firewall Azure/Nama DB):", err)
	} else {
		log.Println("Berhasil nyambung ke Database Azure MySQL!")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	DB = db
}
