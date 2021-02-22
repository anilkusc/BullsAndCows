package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/proullon/ramsql/driver"
)

var a App

func TestReadUserHandler(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{

		{input: "test", output: "hello"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/ReadUser", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.ReadUserHandler)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != test.output {
			t.Errorf("Response is: %v . Expected: %v", string(body), test.output)
		}
	}
}
func TestCreateGameHandler(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{

		{input: "test", output: "hello"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/CreateGame", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.CreateGameHandler)

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

		{input: "test", output: "hello"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/JoinGame", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.JoinGameHandler)

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

		{input: "test", output: "hello"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/StartGame", nil)
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
		req, err := http.NewRequest("POST", "/backend/MakePrediction", nil)
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
		req, err := http.NewRequest("POST", "/backend/Connect", nil)
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
