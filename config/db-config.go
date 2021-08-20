package config

import (
	"database/sql"
	"os"
        _ "github.com/lib/pq"
)

func DbConn() (db *sql.DB) {
	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		dbDriver = "mysql"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		dbPass = ""
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "contact"
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = dbUser+":"+dbPass+"@/"+dbName
	}

	db, err := sql.Open(dbDriver, dbUrl)
	if err != nil {
		panic(err.Error())
	}
	return db
}
