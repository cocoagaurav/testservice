package handler

import "net/http"

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "sessiontoken",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
