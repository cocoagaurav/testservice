package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/firebase"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"github.com/cocoagaurav/httpHandler/model"
	"net/http"
	"time"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlPages.Formpage)
}
func Loginhandler(w http.ResponseWriter, r *http.Request) {

	var authId string
	Db := r.Context().Value("database").(*sql.DB)

	loginUser := &model.User{}

	err := json.NewDecoder(r.Body).Decode(loginUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//fmt.Printf("loginUser :[%+v]", loginUser)

	cred := Db.QueryRow("select auth_id "+
		"						from user "+
		"						where email_id=? AND password=?", loginUser.EmailId, loginUser.Password)

	err = cred.Scan(&authId)

	//fmt.Println("database values are:", name, age, authId)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	token := firebase.GenerateToken(authId)
	http.SetCookie(w, &http.Cookie{
		Name:    "sessiontoken",
		Value:   authId,
		Expires: time.Now().Add(24 * time.Hour),
	})

	json.NewEncoder(w).Encode(token)

}
