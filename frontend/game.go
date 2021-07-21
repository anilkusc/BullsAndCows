package main

import (
	"syscall/js"
	"fmt"
	"encoding/json"
	//"strconv"
)
var (
	window = js.Global()
	doc = window.Get("document")
	body = doc.Get("body")
)

func GetData() {
	var response map[string]interface{}
	json.Unmarshal([]byte(window.Get("localStorage").Get("response").String()), &response)
	session := response["session"].(map[string]interface{})

	turn := doc.Call("getElementById", "turn")
	turn.Set("innerHTML", session["turn"])
	sessionId := doc.Call("getElementById", "session")
	sessionId.Set("innerHTML", session["id"])
	whosturn := doc.Call("getElementById", "whosturn")
	whosturn.Set("innerHTML", session["predictor"])
	player1 := session["player1"].(map[string]interface{})
	player2 := session["player2"].(map[string]interface{})
	players := fmt.Sprintf("%v - %v", player1["name"],player2["name"])
	whosplaying := doc.Call("getElementById", "players")
	whosplaying.Set("innerHTML", players)
	//body.Call("appendChild", turn)
}

//func registerCallbacks() {
//	js.Global().Set("Test", js.FuncOf(Test))
//}

func main() {
	c := make(chan bool)
	GetData()
	<-c
}