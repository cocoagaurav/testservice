package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/olivere/elastic"
	"time"
)

func ElasticConn() *elastic.Client {
	ElasticClient, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		log.Printf("err=[%v]", err)
		time.Sleep(5 * time.Second)
		ElasticConn()
	}
	fmt.Printf("Starting elastic coonection \n\n")
	return ElasticClient
}
