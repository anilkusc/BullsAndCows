package main

import (
	"io"
	"net/http"
)

// ReadUser method reads the user and return it as a json
func (a *App) ReadUserHandler(w http.ResponseWriter, r *http.Request) {
	returnValue := "hello"
	io.WriteString(w, string(returnValue))
	return
}

// CreateGame method creates new game.
func (a *App) CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	returnValue := "hello"
	io.WriteString(w, string(returnValue))
	return
}

// JoinGame method add player to a created game.
func (a *App) JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	returnValue := "hello"
	io.WriteString(w, string(returnValue))
	return
}

// StartGame method is starts the game after both of the players ready
func (a *App) StartGameHandler(w http.ResponseWriter, r *http.Request) {
	returnValue := "hello"
	io.WriteString(w, string(returnValue))
	return
}

// MakePrediction method is in charge for making prediction about opponent player's number.
func (a *App) MakePredictionHandler(w http.ResponseWriter, r *http.Request) {
	returnValue := "hello"
	io.WriteString(w, string(returnValue))
	return
}

// Connect method will be websocket that notify to user for turn
func (a *App) ConnectHandler(w http.ResponseWriter, r *http.Request) {
	returnValue := "hello"
	io.WriteString(w, string(returnValue))
	return
}
