package main

import (
	"database/sql"
	"log"
	"time"
)

var Database *sql.DB

func Opendatabase() *sql.DB {
	var err error
	log.Printf("database is connecting in func.......")
	Database, err = sql.Open("mysql", "root:password123@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("not able to connect to database")
		time.Sleep(5 * time.Second)
		Opendatabase()
	}
	log.Printf("database is connected.......")

	return Database

}
