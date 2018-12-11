package main

import (
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
	"time"
)

var Conn *amqp.Connection

func RabbitConn() {
	var err error
	Conn, err = amqp.Dial("amqp://guest:guest@rabbitmq-server:5672/")
	if err != nil {
		log.Printf("not able to connect rabbitmq")
		time.Sleep(5 * time.Second)
		RabbitConn()
	}
	log.Printf("rabbitmq is connected\n\n\n")

}
