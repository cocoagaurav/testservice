package main

import (
	"fmt"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/streadway/amqp"
	"log"
	"time"
)

var Connection *amqp.Connection

func RabbitConn(env model.Env) *amqp.Connection {
	var err error
	Connection, err = amqp.Dial(env.RabbitUrl)
	if err != nil {
		log.Printf("not able to connect to rabbitmq")
		time.Sleep(5 * time.Second)
		RabbitConn(env)
	}
	fmt.Printf("connected to rabbitmq...... \n\n")
	return Connection
}
