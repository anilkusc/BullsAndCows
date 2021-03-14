package main

import (
	//"encoding/json"
	"io"
	//"log"
	"net/http"
	//"time"
	//"strconv"
	"github.com/anilkusc/BullsAndCows/database"
	//"github.com/anilkusc/BullsAndCows/models"
)

type User struct {
	*database.User
}
type Session struct {
	*database.Session
}
type Move struct {
	*database.Move
}

var u User
var s Session
var m Move

// CreateGame method creates new game.
func (a *App) CreateGameHandler(w http.ResponseWriter, r *http.Request) {
//TODO: check if game started and smt like that
/*	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error decoding user")
		io.WriteString(w, `{"error":"Error decoding user"}`)
		return
	}
	user, err = u.CreateUser(a.DB, user)
	if err != nil {
		log.Println("Error creating user")
		io.WriteString(w, `{"error":"Error creating user"}`)
		return
	}
	var session models.Session
	session.Date = time.Now().Format("02-Jan-2006")
	session.Player1 = user
	session, err = s.CreateSession(a.DB, session)
	if err != nil {
		log.Println("Error creating session")
		io.WriteString(w, `{"error":"Error creating session"}`)
		return
	}
	move := models.Move{
		Session:       session,
		Clue:          models.Clue{Positive: 0, Negative: 0},
		Predictor:     0,
		Prediction:    0,
		Action:        "Created",
	}
	move, err = m.CreateMove(a.DB, move)
	if err != nil {
		log.Println("Error creating move")
		io.WriteString(w, `{"error":"Error creating move"}`)
		return
	}

	returnValue, err := json.Marshal(move)
	if err != nil {
		log.Println("Error marshalling move")
		io.WriteString(w, `{"error":"Error marshalling move"}`)
		return
	}
	io.WriteString(w, string(returnValue))
	return*/
	returnValue := "hello"
	io.WriteString(w, string(returnValue))
	return

}

// JoinGame method add player to a created game.
func (a *App) JoinGameHandler(w http.ResponseWriter, r *http.Request) {
//TODO: check if game started and smt like that
/*	type JoinGame struct {
		Session models.Session `json:"session"`
		User models.User `json:"user"`
	}
	
	var joinGame JoinGame
	err := json.NewDecoder(r.Body).Decode(&joinGame)
	if err != nil {
		log.Println("Error decoding joinGame")
		io.WriteString(w, `{"error":"Error decoding joinGame"}`)
		return
	}
	session, err := s.ReadSession(a.DB, joinGame.Session.Id)
	if err != nil {
		log.Println("There is no such a session with id: ",strconv.Itoa(joinGame.Session.Id))
		io.WriteString(w, `{"error":"There is no such a session with id: `+strconv.Itoa(joinGame.Session.Id)+`"}`)
		return
	}

	moves,err := m.ListMoves(a.DB,session.Id)
	if err != nil {
		log.Println("Cannot List Moves")
		io.WriteString(w, `{"error":"Cannot List Moves"}`)
		return
	}
	move := moves[len(moves)-1]
	move.Action = "Joined"
	move.Session = session
	move.Player2 = joinGame.User
	move.Id++
	move, err = m.CreateMove(a.DB, move)
	if err != nil {
		log.Println("Error creating move")
		io.WriteString(w, `{"error":"Error creating move"}`)
		return
	}

	returnValue, err := json.Marshal(move)
	if err != nil {
		log.Println("Error marshalling move")
		io.WriteString(w, `{"error":"Error marshalling move"}`)
		return
	}
	io.WriteString(w, string(returnValue))
	return*/
	returnValue := "hello"
	io.WriteString(w, string(returnValue))
	return
}

// GetReadyHandler method is starts the game after both of the players ready
func (a *App) GetReadyHandler(w http.ResponseWriter, r *http.Request) {
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
