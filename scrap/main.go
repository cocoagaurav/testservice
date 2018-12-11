package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
	"net/http"
)

func main() {
	RabbitConn()
	ch, err := Conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Printf("error while creating channel")
		return
	}
	_, err = ch.QueueDeclare(
		"quoteQ",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Printf("error while creating a Q")
	}

	dates, err := ch.Consume(
		"quoteQ",
		"",
		false,
		false,
		false,
		false,
		nil)

	if err != nil {
		log.Fatal(err)
		return
	}
	//	scrap := make(chan bool)
	go func() {
		for date := range dates {
			resp := Scrap(string(date.Body))
			err = ch.Publish(
				"",
				date.ReplyTo,
				false,
				false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: date.CorrelationId,
					Body:          []byte(resp),
				})
			date.Ack(false)
		}
	}()
	fmt.Println("scrapping....")
	http.ListenAndServe(":8082", nil)
	//	<-scrap
}
