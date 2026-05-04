package config
import (
	"database/sql"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB

func ConnectDB() {
	dsn := "root:nA^94eV1sha@tcp(127.0.0.1:3307)/IMPAL?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal connect database:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Database tidak merespon:", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)
	DB = db
	log.Println("Database connected")
}
