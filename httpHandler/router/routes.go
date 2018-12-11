package router

import (
	"github.com/cocoagaurav/httpHandler/handler"
	"github.com/cocoagaurav/httpHandler/middleware"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/go-chi/chi"
	"net/http"
)

func Setuproutes(config *model.Configs) http.Handler {
	var route = chi.NewRouter()

	route.Use(middleware.DatabaseMiddleWare(config))

	route.Post("/", handler.FormHandler)
	route.Post("/login", handler.Loginhandler)
	route.Post("/registerform", handler.RegisterformHandler)
	route.Post("/register", handler.RegisterHandler)

	route.Mount("/", Authroutes(config))

	return route

}

func Authroutes(config *model.Configs) http.Handler {

	var route = chi.NewRouter()

	route.Use(middleware.DatabaseMiddleWare(config))
	route.Use(middleware.ElasticMiddleWare(config))
	route.Use(middleware.RabbitMiddleWare(config))
	route.Use(middleware.UserMiddleware(config.Db))

	route.Post("/post", handler.Posthandler)
	route.Post("/logout", handler.LogoutHandler)
	route.Post("/fetchformhandler", handler.Fetchformhandler)
	route.Post("/fetch", handler.FetchHandler)
	route.Get("/quote/{date}", handler.Getquote)

	return route

	//http.Handle("/", route)

}
