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
	
	submitbutton := doc.Call("getElementById", "submitbutton")
	submitbutton.Set("disabled", true)
	abandonbutton := doc.Call("getElementById", "abandonbutton")
	abandonbutton.Set("disabled", true)
	predictionbar := doc.Call("getElementById", "predictionbar")
	predictionbar.Set("disabled", true)
	var title string
	switch fmt.Sprintf("%v",session["start"]) {
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
		predictionbar.Set("disabled", false)
		submitbutton.Set("disabled", true)
		abandonbutton.Set("disabled", true)
		break
	default:
		title = "Error"
		break
	}	
	predictiontitle := doc.Call("getElementById", "predictiontitle")
	predictiontitle.Set("innerHTML", title)
	CreateTable()
	
	//body.Call("appendChild", turn)
}

func CreateTable() {
	historytablebody := doc.Call("getElementById", "historytablebody")
	tr := doc.Call("createElement", "tr")
	td := doc.Call("createElement", "td")
	td2 := doc.Call("createElement", "td")
	td3 := doc.Call("createElement", "td")
	td.Set("innerHTML","deneme")
	td2.Set("innerHTML","deneme")
	td3.Set("innerHTML","deneme")
	historytablebody.Call("appendChild", tr)
	tr.Call("appendChild", td)
	tr.Call("appendChild", td2)
	tr.Call("appendChild", td3)
	/*tr := historytablebody.Call("createElement", "tr")
	historytablebody.Call("appendChild", tr)
	td := tr.Call("createElement", "td")
	td.Set("innerHTML", "message")
	tr.Call("appendChild", td)*/
		
}


//func registerCallbacks() {
//	js.Global().Set("Test", js.FuncOf(Test))
//}

func main() {
	c := make(chan bool)
	GetData()
	<-c
}