package main

import (
	"bytes"
	"context"
	"github.com/cocoagaurav/httpHandler/firebase"
	"github.com/cocoagaurav/httpHandler/handler"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/go-chi/chi"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTestproc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testproc Suite")
}

var _ = Describe("test the register handler", func() {
	var env model.Env
	err := envconfig.Process("myapi", &env)
	if err != nil {
		log.Printf("error while getting env variables:%s", err.Error())
		return
	}
	firebase.FirebaseStartAuth(env)
	r := chi.NewRouter()
	Db := Opendatabase(env)
	MigrateUp(Db)
	It("will run the register handler", func() {
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(`{"emailid":"gaurav@api.com","password":"simple","name":"gaurav","age":23}`)))
		ctx := context.WithValue(req.Context(), "database", Db)
		rr := httptest.NewRecorder()
		r.HandleFunc("/register", handler.RegisterHandler)
		r.ServeHTTP(rr, req.WithContext(ctx))
		Expect(err).ShouldNot(HaveOccurred())
		Expect(rr.Code).To(Equal(http.StatusCreated))
	})

	It("will run the register handler for existing user", func() {
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(`{"emailid":"bharadwaj@api.com","password":"simple","name":"bharadwaj","age":23}`)))
		ctx := context.WithValue(req.Context(), "database", Db)
		rr := httptest.NewRecorder()
		r.HandleFunc("/register", handler.RegisterHandler)
		r.ServeHTTP(rr, req.WithContext(ctx))
		Expect(err).ShouldNot(HaveOccurred())
		Expect(rr.Code).To(Equal(400))
	})

	Describe("test the login handler", func() {

		It("will test login handler for wrong user id", func() {
			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{"emailid":"fakeUser@api.com","password":"simple"}`)))
			rr := httptest.NewRecorder()
			ctx := context.WithValue(req.Context(), "database", Db)

			r.HandleFunc("/login", handler.Loginhandler)
			r.ServeHTTP(rr, req.WithContext(ctx))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(rr.Code).To(Equal(http.StatusNonAuthoritativeInfo))

		})

		It("will test login handler for right user", func() {
			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{"emailid":"gaurav@api.com","password":"simple"}`)))
			ctx := context.WithValue(req.Context(), "database", Db)

			rr := httptest.NewRecorder()
			r.HandleFunc("/login", handler.Loginhandler)
			r.ServeHTTP(rr, req.WithContext(ctx))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(rr.Code).To(Equal(http.StatusOK))

		})
	})
	Describe("remove the new user", func() {
		It("remove user from firebase", func() {
			var uid string
			email := "gaurav@api.com"
			err := Db.QueryRow("select auth_id from user where email_id = ?", email).Scan(&uid)
			firebase.DeleteFirebaseUser(uid)
			Expect(err).ShouldNot(HaveOccurred())

		})
		It("remove user from database", func() {
			q, err := Db.Prepare("delete from user where name = 'gaurav' ")
			q.Exec()
			Expect(err).ShouldNot(HaveOccurred())
		})

	})

})
