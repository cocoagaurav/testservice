package middleware

import (
	"context"
	"database/sql"
	"github.com/cocoagaurav/httpHandler/model"
	"log"
	"net/http"
)

func UserMiddleware(Db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var User model.User
			c, err := r.Cookie("sessiontoken")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = Db.QueryRow("SELECT name,email_id   "+
				"				FROM user"+
				"				WHERE auth_id = ?", c.Value).Scan(&User.Name, &User.EmailId)

			if err != nil {
				log.Printf("error while getting user info :%v", err)
			}

			if User.Name == "" && User.EmailId == "" {
				w.WriteHeader(http.StatusNoContent)
			} else {
				ctx := context.WithValue(r.Context(), "user", User)
				next.ServeHTTP(w, r.WithContext(ctx))
			}

		})
	}
}
func DatabaseMiddleWare(config *model.Configs) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "database", config.Db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ElasticMiddleWare(config *model.Configs) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "elastic", config.Ec)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RabbitMiddleWare(config *model.Configs) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "rabbit", config.Rabbit)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
