package handler

import (
	"fmt"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/go-chi/chi"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
	"net/http"
)

func Getquote(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "date")
	Conn := r.Context().Value("rabbit").(*amqp.Connection)
	User := r.Context().Value("user").(model.User)
	ch, err := Conn.Channel()
	if err != nil {
		log.Printf("error while creating channle")
		return
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"response",
		false,
		false,
		false,
		false,
		nil)

	_, err = ch.QueueDeclare(
		"quoteQ",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Printf("error while creating a Q:%v", err)
	}
	err = ch.Publish(
		"",
		"quoteQ",
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			ReplyTo:       "response",
			CorrelationId: User.EmailId,
			Body:          []byte(params),
		})

	msg, err := ch.Consume(
		"response",
		"",
		false,
		false,
		false,
		false,
		nil)

	finish := make(chan bool)

	go func() {
		for mssg := range msg {
			if User.EmailId == mssg.CorrelationId {
				fmt.Fprint(w, string(mssg.Body))
				mssg.Ack(false)
				break

			}
		}
		finish <- true

	}()
	<-finish
}
