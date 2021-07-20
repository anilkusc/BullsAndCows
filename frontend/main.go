package main

import (
	"strconv"
	"syscall/js"
	"log"
	"net/http"
	"bytes"
    "encoding/json"
	"io/ioutil"
)
func request(URL string,payload string){

	//values := map[string]string{"name": "John Doe", "occupation": "gardener"}
    json_data, err := json.Marshal(payload)

    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post(URL, "application/json", bytes.NewBuffer(json_data))

    if err != nil {
        log.Fatal(err)
    }
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
func operate(this js.Value, inputs []js.Value) interface{} {
	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")
	total := doc.Call("getElementById", "total")
	go request("https://httpbin.org/post","{'name': 'John Doe', 'occupation': 'gardener'}")
	total.Set("innerHTML", strconv.Itoa(3))
	body.Call("appendChild", total)
	return nil
}

func registerCallbacks() {
	js.Global().Set("operate", js.FuncOf(operate))
}

func main() {
	c := make(chan bool)
	registerCallbacks()
	<-c
}