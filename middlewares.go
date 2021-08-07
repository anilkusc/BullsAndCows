package main

import (
	"net/http"
	"strconv"
)

func Inbound(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type,Secret,User,Session")
		w.Header().Set("Access-Control-Expose-Headers", "Secret,User,Session")
		next(w, r)
	}
}
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			return
		}
		client_secret, _ := strconv.Atoi(r.Header.Get("Secret"))
		user, _ := strconv.Atoi(r.Header.Get("User"))
		session, _ := strconv.Atoi(r.Header.Get("Session"))
		secret, exist := store[session][user]
		if exist {
			if secret == client_secret {
				next(w, r)
			}
		}
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
}
