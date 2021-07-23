package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"

	//"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/anilkusc/BullsAndCows/database"
	"github.com/anilkusc/BullsAndCows/logic"
	"github.com/anilkusc/BullsAndCows/models"
	"github.com/gorilla/websocket"
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

	bodyString, _ := ioutil.ReadAll(r.Body)
	var user models.User

	err := json.Unmarshal([]byte(string(bodyString)), &user)
	if err != nil {
		log.Println("Error decoding user", err)
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
		Session:    session,
		Clue:       models.Clue{Positive: 0, Negative: 0},
		Prediction: 0,
		Action:     "Created",
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
	return
}

// JoinGame method add player to a created game.
func (a *App) JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: check if game started and smt like that
	type JoinGame struct {
		Session models.Session `json:"session"`
		User    models.User    `json:"user"`
	}
	bodyString, _ := ioutil.ReadAll(r.Body)
	var joinGame JoinGame

	err := json.Unmarshal([]byte(string(bodyString)), &joinGame)
	if err != nil {
		log.Println("Error decoding joinGame", err)
		io.WriteString(w, `{"error":"Error decoding joinGame"}`)
		return
	}

	session, err := s.ReadSession(a.DB, joinGame.Session.Id)
	if err != nil {
		log.Println("There is no such a session with id: ", strconv.Itoa(joinGame.Session.Id))
		io.WriteString(w, `{"error":"There is no such a session with id: `+strconv.Itoa(joinGame.Session.Id)+`"}`)
		return
	}

	if session.Player2.Name == "" {
		user, err := u.CreateUser(a.DB, joinGame.User)
		if err != nil {
			log.Println("Error creating user")
			io.WriteString(w, `{"error":"Error creating user."`)
			return
		}
		session.Player2 = user
		session, err = s.UpdateSession(a.DB, session)
		if err != nil {
			log.Println("Cannot update player 2")
			io.WriteString(w, `{"error":"Cannot update player 2"`)
			return
		}
	} else {
		if session.Player1.Name != joinGame.User.Name && session.Player2.Name != joinGame.User.Name {
			log.Println("Wrong Username for Joining")
			io.WriteString(w, `{"error":"Wrong Username for Joining"`)
			return
		}
	}

	move := models.Move{
		Session: session,
		Action:  "Joined",
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
	return
}

// GetReadyHandler method is starts the game after both of the players ready
func (a *App) GetReadyHandler(w http.ResponseWriter, r *http.Request) {
	type GetReady struct {
		Number  int `json:"number"`
		User    int `json:"user"`
		Session int `json:"session"`
	}

	bodyString, _ := ioutil.ReadAll(r.Body)
	var getReady GetReady

	err := json.Unmarshal([]byte(string(bodyString)), &getReady)
	if err != nil {
		log.Println("Error decoding getReady")
		io.WriteString(w, `{"error":"Error decoding getReady"}`)
		return
	}

	if getReady.User >= 3 || getReady.User <= 0 {
		log.Println("Neither Player 1 Nor Player 2")
		io.WriteString(w, `{"error":"Neither Player 1 Nor Player 2"}`)
		return
	}

	_, err = u.ReadUser(a.DB, getReady.User)
	if err != nil {
		log.Println("Error getting user")
		io.WriteString(w, `{"error":"Error getting user"}`)
		return
	}
	session, err := s.ReadSession(a.DB, getReady.Session)
	if err != nil {
		log.Println("Error getting session")
		io.WriteString(w, `{"error":"Error getting session"}`)
		return
	}
	var action string
	if getReady.User == 1 && session.Start == 1 || getReady.User == 2 && session.Start == 2 {
		log.Println("You've already get ready")
		io.WriteString(w, `{"error":"You've already get ready"}`)
		return
	} else {
		if getReady.User == 1 {
			session.Player1Number = getReady.Number
			if session.Start == 2 {
				session.Start = 3
				action = "Started"
				session.Predictor = 1
				session.Turn = 1
			} else {
				session.Start = 1
				action = "Ready1"
			}
		} else {
			session.Player2Number = getReady.Number
			if session.Start == 1 {
				session.Start = 3
				action = "Started"
				session.Predictor = 1
				session.Turn = 1
			} else {
				session.Start = 2
				action = "Ready2"
			}
		}

	}

	session, err = s.UpdateSession(a.DB, session)
	if err != nil {
		log.Println("Error updating session")
		io.WriteString(w, `{"error":"Error updating session"}`)
		return
	}

	move := models.Move{
		Session: session,
		Action:  action,
	}
	move, err = m.CreateMove(a.DB, move)
	if err != nil {
		log.Println("Error creating move")
		io.WriteString(w, `{"error":"Error creating move"}`)
		return
	}

	if getReady.User == 1 {
		move.Session.Player2Number = 0
	} else {
		move.Session.Player1Number = 0
	}

	returnValue, err := json.Marshal(move)
	if err != nil {
		log.Println("Error marshalling move")
		io.WriteString(w, `{"error":"Error marshalling move"}`)
		return
	}
	io.WriteString(w, string(returnValue))
	return
}

// Connect method will be websocket that notify to user for turn
func (a *App) ConnectHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Cannot create Websocket")
		io.WriteString(w, `{"error":"Cannot create Websocket"}`)
		return
	}
	for {
		messageType, req, _ := ws.ReadMessage()
		for {
			type Connect struct {
				User    int `json:"user"`
				Session int `json:"session"`
			}
			var connect Connect
			if err := json.Unmarshal(req, &connect); err != nil {
				log.Println("Cannot unmarshall")
			}

			moves, err := m.ListMoves(a.DB, connect.Session)
			if err != nil {
				log.Println("Cannot list moves")
			}
			for _, move := range moves {
				move.Session.Player1Number = 0
				move.Session.Player2Number = 0
			}
			//move := moves[len(moves)-1]

			returnValue, err := json.Marshal(moves)
			if err != nil {
				log.Println("Error marshalling move")
				io.WriteString(w, `{"error":"Error marshalling move"}`)
				return
			}

			data := []byte(returnValue)
			if err := ws.WriteMessage(messageType, data); err != nil {
				log.Println("error while sending message:", err)
				return
			}
			time.Sleep(5 * time.Second)
		}
	}
}

// MakePrediction method is in charge for making prediction about opponent player's number.
func (a *App) MakePredictionHandler(w http.ResponseWriter, r *http.Request) {
	type Prediction struct {
		Number  int `json:"prediction"`
		User    int `json:"user"`
		Session int `json:"session"`
	}
	var prediction Prediction
	err := json.NewDecoder(r.Body).Decode(&prediction)
	if err != nil {
		log.Println("Error decoding prediction")
		io.WriteString(w, `{"error":"Error decoding prediction"}`)
		return
	}
	user, err := u.ReadUser(a.DB, prediction.User)
	if err != nil {
		log.Println("Error getting user")
		io.WriteString(w, `{"error":"Error getting user"}`)
		return
	}
	session, err := s.ReadSession(a.DB, prediction.Session)
	if err != nil {
		log.Println("Error getting session")
		io.WriteString(w, `{"error":"Error getting session"}`)
		return
	}

	var clue models.Clue
	var action string

	if session.Player1 != user && session.Player2 != user {
		log.Println("There is no user in this session")
		io.WriteString(w, `{"error":"There is no user in this session"}`)
		return
	} else {
		if session.Player1 == user {
			if session.Predictor != 1 {
				log.Println("It is not your turn!")
				io.WriteString(w, `{"error":"It is not your turn!"}`)
				return
			}
			clue, err = logic.CalculateClue(prediction.Number, session.Player2Number)
			if err != nil {
				log.Println("Error calculating clue")
				io.WriteString(w, `{"error":"Error calculating clue"}`)
				return
			}
			if clue.Negative == 100 && clue.Positive == 100 {
				session.Winner = 1
				session.End = 1
				action = "End"
			} else {
				action = "Predicted"
			}
		} else {
			if session.Predictor != 2 {
				log.Println("It is not your turn!")
				io.WriteString(w, `{"error":"It is not your turn!"}`)
				return
			}
			clue, _ = logic.CalculateClue(prediction.Number, session.Player1Number)
			if err != nil {
				log.Println("Error calculating clue")
				io.WriteString(w, `{"error":"Error calculating clue"}`)
				return
			}
			if clue.Negative == 100 && clue.Positive == 100 {
				session.Winner = 2
				session.End = 2
				action = "End"
			} else {
				action = "Predicted"
			}
		}
	}

	moves, err := m.ListMoves(a.DB, session.Id)
	if err != nil {
		log.Println("Cannot list moves")
		io.WriteString(w, `{"error":"Cannot list moves"}`)
		return
	}
	move := moves[len(moves)-1]
	move.Session = session
	move.Clue = clue
	move.Prediction = prediction.Number
	move.Action = action

	move, err = m.CreateMove(a.DB, move)
	if err != nil {
		log.Println("Cannot create move")
		io.WriteString(w, `{"error":"Cannot create move"}`)
		return
	}

	session.Turn++
	if session.Predictor == 1 {
		session.Predictor = 2
	} else {
		session.Predictor = 1
	}
	session, err = s.UpdateSession(a.DB, session)
	if err != nil {
		log.Println("Error updating session")
		io.WriteString(w, `{"error":"Error updating session"}`)
		return
	}
	if prediction.User == 1 {
		move.Session.Player2Number = 0
	} else {
		move.Session.Player1Number = 0
	}
	returnValue, err := json.Marshal(move)
	if err != nil {
		log.Println("Error marshalling move")
		io.WriteString(w, `{"error":"Error marshalling move"}`)
		return
	}
	io.WriteString(w, string(returnValue))
	return
}
