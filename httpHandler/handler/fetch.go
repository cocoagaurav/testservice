package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/labstack/gommon/log"
	"net/http"
)

func Fetchformhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlPages.Fetchform)
}

func FetchHandler(w http.ResponseWriter, r *http.Request) {
	Db := r.Context().Value("database").(*sql.DB)
	userpost := &model.User{}
	post := &model.Post{}
	//ElasticClient := r.Context().Value("elastic").(*elastic.Client)
	err := json.NewDecoder(r.Body).Decode(userpost)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	}
	rows, err := Db.Query("SELECT * FROM post WHERE id = ?", userpost.EmailId)
	defer rows.Close()
	if err != nil {
		log.Printf("error while getting post:%v", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	for rows.Next() {
		err := rows.Scan(&post.Name, &post.EmailId, &post.Title, &post.Discription)
		if err != nil {
			log.Printf("error while scaning rows:%v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		post := fmt.Sprintf("Name = %s   Title = %s   Discription = %s", post.Name, post.Title, post.Discription)

		json.NewEncoder(w).Encode(post)
	}
	//esquery := elastic.NewTermQuery("id", userpost.AuthId)
	//result, err := ElasticClient.Search("userpost").Index("userpost").Type("post").Query(esquery).Do(context.Background())
	//if err != nil {
	//	log.Printf("error is: [%v]", err.Error())
	//}
	//for _, hit := range result.Hits.Hits {
	//	json.Unmarshal(*hit.Source, post)
	//	fmt.Fprintf(w, "USERID=%d \n\n title=%s \n\n description=%s \n\n\n\n ", userpost.AuthId, post.Title, post.Discription)
	//}
	//http.Redirect(w, r, "/fetchformhandler", 302)
}
