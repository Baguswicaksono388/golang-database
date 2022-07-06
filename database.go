package golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/bisnishub_auth")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  // awal koneksi membuat 10 koneksi
	db.SetMaxOpenConns(100)                 // maksimal yg dibuat 100 koneksi
	db.SetConnMaxIdleTime(5 * time.Minute)  // koneksi tidak aktif lebih dari 5 menit
	db.SetConnMaxLifetime(60 * time.Minute) // akan memperbarui koneksi setelah 60 menit

	return db
}
