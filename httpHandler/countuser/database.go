package main

import (
	"database/sql"
	"fmt"
	"time"
)

func Openbase() *sql.DB {
	Db, err := sql.Open("mysql", "root:password123@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {

		fmt.Printf("not able to connect to database:%v", err)
		time.Sleep(4 * time.Second)
		Openbase()
	}
	fmt.Printf("Starting database... \n\n")
	return Db
}
