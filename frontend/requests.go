package main

import (
	"log"
	"net/http"
	"bytes"
    "encoding/json"
	"io/ioutil"
)

func Request(URL string,payload string){

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