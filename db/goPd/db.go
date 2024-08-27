package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "esbon"
	password = "Esbon@0925"
	dbname   = "godb"
)

var DB *sql.DB

func OpenDatabase() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Successfully connected!")
	return nil

}

func CloseDatabase() error {
	return DB.Close()
}
