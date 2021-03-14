package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	_ "github.com/proullon/ramsql/driver"
)

var now = time.Now().Format("02-Jan-2006")

func TestCreateGameHandler(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{

		{input: `test`, output: "hello"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/CreateGame",  strings.NewReader(test.input))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.StartGameHandler)

		handler.ServeHTTP(rr, req)

		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != test.output {
			t.Errorf("Response is: %v . Expected: %v", string(body), test.output)
		}
	}
}
func TestGetReadyHandler(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{

		{input: `test`, output: "hello"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/GetReady",  strings.NewReader(test.input))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.StartGameHandler)

		handler.ServeHTTP(rr, req)

		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != test.output {
			t.Errorf("Response is: %v . Expected: %v", string(body), test.output)
		}
	}
}

func TestJoinGameHandler(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{

		{input: `test`, output: "hello"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/JoinGame",  strings.NewReader(test.input))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.StartGameHandler)

		handler.ServeHTTP(rr, req)

		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != test.output {
			t.Errorf("Response is: %v . Expected: %v", string(body), test.output)
		}
	}
}
func TestStartGameHandler(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{

		{input: `test`, output: "hello"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/StartGame",  strings.NewReader(test.input))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.StartGameHandler)

		handler.ServeHTTP(rr, req)

		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != test.output {
			t.Errorf("Response is: %v . Expected: %v", string(body), test.output)
		}
	}
}
func TestMakePredictionHandler(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{

		{input: "test", output: "hello"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/MakePrediction",  strings.NewReader(test.input))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.MakePredictionHandler)

		handler.ServeHTTP(rr, req)

		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != test.output {
			t.Errorf("Response is: %v . Expected: %v", string(body), test.output)
		}
	}
}
func TestConnectHandler(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{

		{input: "test", output: "hello"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/Connect",  strings.NewReader(test.input))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.ConnectHandler)

		handler.ServeHTTP(rr, req)

		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != test.output {
			t.Errorf("Response is: %v . Expected: %v", string(body), test.output)
		}
	}
}
