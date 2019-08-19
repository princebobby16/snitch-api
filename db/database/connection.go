package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "snitch"
	password = "snitch"
	dbname   = "incidentReport"
)

var DBConn *sql.DB

func Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	DBConn = db
	return nil
}

func Disconnect() error {
	err := DBConn.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
