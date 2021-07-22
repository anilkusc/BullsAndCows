package main

import (
	"log"
	"net/http"

	"strings"
	"syscall/js"
)

var (
	window = js.Global()
	doc    = window.Get("document")
	body   = doc.Get("body")
)

func GetReady(this js.Value, inputs []js.Value) interface{} {
	go func() {

		body := strings.NewReader("{ \"user\": " + window.Get("localStorage").Get("user").String() + " , \"session\": " + window.Get("localStorage").Get("session").String() + "  ,\"number\": " + inputs[0].String() + " }")
		//body := "{ 'user': { 'name': '" + window.Get("localStorage").Get("username").String() + "' }, 'session': { 'id': " + window.Get("localStorage").Get("session").String() + " }, 'number': " + inputs[0].String() + " }"
		log.Println(body)
		req, err := http.NewRequest("POST", "http://localhost:8080/backend/GetReady", body)
		if err != nil {
			log.Println(err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
	}()
	return nil
}

func GetData() {
	//var response Move
	//json.Unmarshal([]byte(window.Get("localStorage").Get("response").String()), &response)

	turn := doc.Call("getElementById", "turn")
	turn.Set("innerHTML", window.Get("localStorage").Get("turn").String())
	sessionId := doc.Call("getElementById", "session")
	sessionId.Set("innerHTML", window.Get("localStorage").Get("session").String())
	whosturn := doc.Call("getElementById", "whosturn")
	whosturn.Set("innerHTML", window.Get("localStorage").Get("whosturn").String())
	whosplaying := doc.Call("getElementById", "players")
	whosplaying.Set("innerHTML", window.Get("localStorage").Get("players").String())

	var title string
	switch window.Get("localStorage").Get("start").String() {
	case "0":
		title = "Waiting for other player..."
		break
	case "1":
		title = "Player1 is Ready"
		break
	case "2":
		title = "Player2 is Ready"
		break
	case "3":
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
}

/*
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
*/
func registerCallbacks() {
	js.Global().Set("GetReady", js.FuncOf(GetReady))
}

func main() {
	c := make(chan bool)
	GetData()
	registerCallbacks()
	<-c
}
