package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/firebase"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/labstack/gommon/log"
	"net/http"
)

func RegisterformHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlPages.Registerpage)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	Db := r.Context().Value("database").(*sql.DB)
	err := Db.Ping()
	if err != nil {
		log.Printf("error while Db.ping err:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	newUser := &model.User{}
	err = json.NewDecoder(r.Body).Decode(newUser)
	if err != nil {
		log.Printf("error while json decoder err:%v", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var uid string
	_ = Db.QueryRow("select auth_id from user where email_id = ?", newUser.EmailId).Scan(&uid)
	if uid != "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cred, err := Db.Prepare("insert into user (name,email_id,password,age,auth_id) value (?,?,?,?,?)")

	defer cred.Close()

	if err != nil {
		log.Printf("error while Db.prepare err:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	user, err := firebase.CreateFireBaseUser(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = cred.Exec(newUser.Name, newUser.EmailId, newUser.Password, newUser.Age, user.UID)
	if (err) != nil {
		log.Printf("error while Db.exec err:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)

}
