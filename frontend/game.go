package main

import (
	"encoding/json"
	"syscall/js"
	//"strconv"
)

var (
	window = js.Global()
	doc    = window.Get("document")
	body   = doc.Get("body")
)

func GetData() {
	var response Move
	json.Unmarshal([]byte(window.Get("localStorage").Get("response").String()), &response)

	turn := doc.Call("getElementById", "turn")
	turn.Set("innerHTML", response.Session.Turn)
	sessionId := doc.Call("getElementById", "session")
	sessionId.Set("innerHTML", response.Session.Id)
	whosturn := doc.Call("getElementById", "whosturn")
	whosturn.Set("innerHTML", response.Session.Predictor)
	players := response.Player1.Name + " - " + response.Player2.Name
	whosplaying := doc.Call("getElementById", "players")
	whosplaying.Set("innerHTML", players)

	var title string
	switch response.Session.Start {
	case 0:
		title = "Waiting for other player..."
		break
	case 1:
		title = "Player1 is Ready"
		break
	case 2:
		title = "Player2 is Ready"
		break
	case 3:
		title = "Prediction"
		submitbutton := doc.Call("getElementById", "submitbutton")
		abandonbutton := doc.Call("getElementById", "abandonbutton")
		readybutton := doc.Call("getElementById", "readybutton")
		predictionbar := doc.Call("getElementById", "predictionbar")
		predictionbar.Set("disabled", false)
		submitbutton.Set("disabled", false)
		readybutton.Set("disabled", true)
		abandonbutton.Set("disabled", false)
		break
	default:
		title = "Error"
		break
	}
	predictiontitle := doc.Call("getElementById", "predictiontitle")
	predictiontitle.Set("innerHTML", title)
	var responses []Move
	responses = append(responses, response)
	CreateTable(responses)
}

func CreateTable(moves []Move) {
	historytablebody := doc.Call("getElementById", "historytablebody")
	for _, move := range moves {
		tr := doc.Call("createElement", "tr")
		historytablebody.Call("appendChild", tr)
		td_id := doc.Call("createElement", "td")
		td_id.Set("innerHTML", move.Id)
		tr.Call("appendChild", td_id)
		td_negative := doc.Call("createElement", "td")
		td_negative.Set("innerHTML", move.Clue.Negative)
		tr.Call("appendChild", td_negative)
		td_positive := doc.Call("createElement", "td")
		td_positive.Set("innerHTML", move.Clue.Positive)
		tr.Call("appendChild", td_positive)
		td_prediction := doc.Call("createElement", "td")
		td_prediction.Set("innerHTML", move.Prediction)
		tr.Call("appendChild", td_prediction)
		td_predictor := doc.Call("createElement", "td")
		td_predictor.Set("innerHTML", move.Session.Predictor)
		tr.Call("appendChild", td_predictor)
	}
}

//func registerCallbacks() {
//	js.Global().Set("Test", js.FuncOf(Test))
//}

func main() {
	c := make(chan bool)
	GetData()
	<-c
}

type Session struct {
	Id            int    `json:"id"`
	Date          string `json:"date"`
	Turn          int    `json:"turn"`          // Turn Count
	Player1       User   `json:"player1"`       // Player1 Name
	Player2       User   `json:"player2"`       // Player2 Name
	Player1Number int    `json:"player1number"` // Player1 Number
	Player2Number int    `json:"player2number"` // Player2 Number
	Predictor     int    `json:"predictor"`     // Who is Predicting
	Start         int    `json:"start"`         // Is Game Started or not.0 is not started.1 is player1 ready.2 is player2 ready.3 is both ready.
	End           int    `json:"end"`           // Is Game ended or not.0 or 1.
	Winner        int    `json:"winner"`        // It indicates the winner.It can be 0(Not Ended),1,2
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Clue struct {
	Positive int `json:"positive"`
	Negative int `json:"negative"`
}

type Move struct {
	Id         int `json:"id"`
	Session    `json:"session"`
	Clue       `json:"clue"` // Clues like +1/-1
	Prediction int           `json:"prediction"` // Prediction
	Action     string        `json:"action"`     // Action
}
