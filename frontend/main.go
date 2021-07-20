package main

import (
	"strconv"
	"syscall/js"
)

func operate(this js.Value, inputs []js.Value) interface{} {
	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")
	total := doc.Call("getElementById", "total")
	go Request("https://httpbin.org/post","{'name': 'John Doe', 'occupation': 'gardener'}")
	total.Set("innerHTML", strconv.Itoa(3))
	body.Call("appendChild", total)
	return nil
}

func CreateGame(this js.Value, inputs []js.Value) interface{} {
	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")
	total := doc.Call("getElementById", "total")
	go Request("http://localhost:8080/backend/CreateGame","{\"name\": \"anil\"}")
	total.Set("innerHTML", strconv.Itoa(30))
	body.Call("appendChild", total)
	return nil
}

func registerCallbacks() {
	js.Global().Set("operate", js.FuncOf(operate))
	js.Global().Set("CreateGame", js.FuncOf(CreateGame))
}

func main() {
	c := make(chan bool)
	registerCallbacks()
	<-c
}