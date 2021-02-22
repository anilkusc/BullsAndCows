package main

import (
	"io"
	"net/http"
)

func (a *App) ReadUser(w http.ResponseWriter, r *http.Request) {
	returnValue := "hello"
	io.WriteString(w, string(returnValue))
	return
}
