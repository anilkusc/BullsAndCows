package main

import (
	"log"
	"net/http"

	"strings"
	"syscall/js"

	"github.com/tidwall/gjson"
)

var (
	window = js.Global()
	doc    = window.Get("document")
	body   = doc.Get("body")
)

func ConnectWebsocket(URL string, message string) {

	ws := window.Get("WebSocket").New(URL)

	ws.Call("addEventListener", "open", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		println("connected to websocket")
		ws.Call("send", message)
		return nil
	}))

	ws.Call("addEventListener", "close", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		/*code := args[0].Get("code").Int()
		println(fmt.Sprintf("websocket close %d\n", code))
		if code == 1000 {
			println("websocket bye!")
		} else {
			go func() {
				select {
				case <-time.After(time.Second * 10):
					connectWebsocket()
				}
			}()
		}*/
		log.Println("Websocket connection is closed")
		return nil
	}))
	ws.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		//p := doc.Call("getElementById", "p")
		//p.Set("innerHTML", args[0].Get("data"))
		CreateTable(args[0].Get("data").String())
		return nil
	}))

	ws.Call("addEventListener", "error", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		/*code := args[0].Get("code").String()
		println(fmt.Sprintf("websocket error %s\n", code))
		if "ECONNREFUSED" == code {
			go func() {
				select {
				case <-time.After(time.Second * 10):
					connectWebsocket()
				}
			}()
		} else {
			println("websocket bye!")
		}*/
		log.Println("websocket message error")
		return nil
	}))
}

func GetReady(this js.Value, inputs []js.Value) interface{} {
	go func() {

		body := strings.NewReader("{ \"user\": " + window.Get("localStorage").Get("user").String() + " , \"session\": " + window.Get("localStorage").Get("session").String() + "  ,\"number\": " + inputs[0].String() + " }")
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
		//a, _ := ioutil.ReadAll(resp.Body)
		//bodyString := string(a)
		submitbutton := doc.Call("getElementById", "submitbutton")
		abandonbutton := doc.Call("getElementById", "abandonbutton")
		readybutton := doc.Call("getElementById", "readybutton")
		predictionbar := doc.Call("getElementById", "predictionbar")
		numberbar := doc.Call("getElementById", "numberbar")
		predictionbar.Set("disabled", false)
		submitbutton.Set("disabled", false)
		readybutton.Set("disabled", true)
		abandonbutton.Set("disabled", false)
		numberbar.Set("disabled", true)
		go ConnectWebsocket("ws://localhost:8080/backend/Connect", "{\"user\": "+window.Get("localStorage").Get("user").String()+", \"session\": "+window.Get("localStorage").Get("session").String()+"}")
		//window.Get("localStorage").Set("a", bodyString)
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

// TODO: create table with websocket responses
func CreateTable(moves string) {
	historytablebody := doc.Call("getElementById", "historytablebody")
	historytablebody.Call("replaceChildren")
	result := gjson.Get(moves, "@this")
	for _, name := range result.Array() {
		tr := doc.Call("createElement", "tr")
		historytablebody.Call("appendChild", tr)
		td_id := doc.Call("createElement", "td")
		td_id.Set("innerHTML", gjson.Get(name.String(), "id").String())
		tr.Call("appendChild", td_id)
		td_negative := doc.Call("createElement", "td")
		td_negative.Set("innerHTML", gjson.Get(name.String(), "clue.negative").String())
		tr.Call("appendChild", td_negative)
		td_positive := doc.Call("createElement", "td")
		td_positive.Set("innerHTML", gjson.Get(name.String(), "clue.positive").String())
		tr.Call("appendChild", td_positive)
		td_prediction := doc.Call("createElement", "td")
		td_prediction.Set("innerHTML", gjson.Get(name.String(), "prediction").String())
		tr.Call("appendChild", td_prediction)
		td_predictor := doc.Call("createElement", "td")
		td_predictor.Set("innerHTML", gjson.Get(name.String(), "session.predictor").String())
		tr.Call("appendChild", td_predictor)

	}
}

func registerCallbacks() {
	js.Global().Set("GetReady", js.FuncOf(GetReady))
}

func main() {
	c := make(chan bool)
	GetData()
	registerCallbacks()
	<-c
}
