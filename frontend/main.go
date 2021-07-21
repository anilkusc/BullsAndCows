package main

import (
	"syscall/js"
	"log"
	"net/http"
	"bytes"
    "encoding/json"
	"io/ioutil"	
)
var (
	window = js.Global()
)

func CreateGameRequest(URL string , userName string) {

	type Payload struct {
		Name string `json:"name"`
	}
	
	payload := Payload{
		Name: userName,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	body := bytes.NewReader(payloadBytes)
	
	req, err := http.NewRequest("POST", URL, body)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	a, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(a)
	
	window.Get("localStorage").Set("response",bodyString)
	window.Set("location","game.html")
	defer resp.Body.Close()
}

func CreateGame(this js.Value, inputs []js.Value) interface{} {
	go CreateGameRequest("http://localhost:8080/backend/CreateGame",inputs[0].String())
	return nil
}

func registerCallbacks() {
	js.Global().Set("CreateGame", js.FuncOf(CreateGame))
}

func main() {
	c := make(chan bool)
	registerCallbacks()
	<-c
}