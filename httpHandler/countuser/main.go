package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	var count int

	db := Openbase()

	err := db.QueryRow("select count(*) from user").Scan(&count)
	if err != nil {
		log.Printf("error while scanning count:%v", err)
	}

	fmt.Printf("total number of users are :%d", count)

	http.ListenAndServe(":8000", nil)

}
