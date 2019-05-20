package config

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)


func DbInit() *sql.DB {
	connStr := fmt.Sprintf("server=%s;Port=%d;User id=%s;Password=%s;database=%s;app name=%s;timeout=30",
		"*****", ***, "**", "****", "******", "***")

	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatal("err")
	}

	if err := db.Ping(); err != nil {
		log.Fatal("err")
	}

	return db
}
