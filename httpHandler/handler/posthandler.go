package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

func AfterLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlPages.InternalPage)
}
func Posthandler(w http.ResponseWriter, r *http.Request) {

	Db := r.Context().Value("database").(*sql.DB)
	Conn := r.Context().Value("rabbit").(*amqp.Connection)
	User := r.Context().Value("user").(model.User)

	err := Db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	newpost := &model.Post{}
	err = json.NewDecoder(r.Body).Decode(newpost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newpost.Name = User.Name
	newpost.EmailId = User.EmailId

	fmt.Printf("\n\npost name:%s \n post title:%s \n post disc:%s ", newpost.Name, newpost.Title, newpost.Discription)

	var uid string
	_ = Db.QueryRow("select auth_id "+
		"					from user "+
		"					where email_id =?", User.EmailId).
		Scan(&uid)

	jsonpost, err := json.Marshal(newpost)
	if err != nil {
		log.Printf("post marshal error:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("json data is:%s", string(jsonpost))

	Ch, err := Conn.Channel()
	if err != nil {
		log.Print(err.Error())
	}
	Q, err := Ch.QueueDeclare(
		"PostQ",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Print(err.Error())
	}
	err = Ch.Publish(
		"",
		Q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonpost,
		})
	if err != nil {
		log.Print(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(newpost)
}
