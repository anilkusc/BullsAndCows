package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"syscall/js"

	"github.com/tidwall/gjson"
	_ "github.com/tidwall/gjson"
)

var (
	window = js.Global()
)

func CreateGameRequest(URI string, userName string) {

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
	req, err := http.NewRequest("POST", "http://"+window.Get("location").Get("hostname").String()+":8080"+URI, body)
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

	window.Get("localStorage").Set("players", gjson.Get(bodyString, "session.player1.name").String()+"-"+gjson.Get(bodyString, "session.player2.name").String())
	window.Get("localStorage").Set("session", gjson.Get(bodyString, "session.id").String())
	window.Get("localStorage").Set("username", gjson.Get(bodyString, "session.player1.name").String())
	window.Get("localStorage").Set("userid", gjson.Get(bodyString, "session.player1.id").String())
	window.Get("localStorage").Set("user", "1")
	window.Set("location", "game.html")
	defer resp.Body.Close()
}

func JoinGameRequest(URI string, sessionId string, userName string) {

	type User struct {
		Name string `json:"name"`
	}
	type Session struct {
		Id int `json:"id"`
	}
	type Payload struct {
		User    `json:"user"`
		Session `json:"session"`
	}
	var user User
	var session Session
	user.Name = userName
	session.Id, _ = strconv.Atoi(sessionId)
	payload := Payload{
		User:    user,
		Session: session,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://"+window.Get("location").Get("hostname").String()+":8080"+URI, body)
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

	window.Get("localStorage").Set("players", gjson.Get(bodyString, "session.player1.name").String()+"-"+gjson.Get(bodyString, "session.player2.name").String())
	window.Get("localStorage").Set("session", gjson.Get(bodyString, "session.id").String())
	window.Get("localStorage").Set("username", gjson.Get(bodyString, "session.player2.name").String())
	window.Get("localStorage").Set("userid", gjson.Get(bodyString, "session.player2.id").String())
	window.Get("localStorage").Set("user", "2")
	window.Set("location", "game.html")
	defer resp.Body.Close()
}

func CreateGame(this js.Value, inputs []js.Value) interface{} {
	go CreateGameRequest("/backend/CreateGame", inputs[0].String())
	return nil
}

func JoinGame(this js.Value, inputs []js.Value) interface{} {
	go JoinGameRequest("/backend/JoinGame", inputs[0].String(), inputs[1].String())
	return nil
}

func registerCallbacks() {
	js.Global().Set("CreateGame", js.FuncOf(CreateGame))
	js.Global().Set("JoinGame", js.FuncOf(JoinGame))
}

func main() {
	c := make(chan bool)
	registerCallbacks()
	<-c
}
