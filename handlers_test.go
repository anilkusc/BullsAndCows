package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"database/sql"
	_ "github.com/proullon/ramsql/driver"
)

var now = time.Now().Format("02-Jan-2006")

func TestCreateGameHandler(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER DEFAULT 0,Negative INTEGER DEFAULT 0,Predictor INTEGER,Prediction INTEGER,Action TEXT);`,
		`CREATE TABLE Users (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);`,
		`CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,Turn INTEGER NOT NULL,Player1Id INTEGER,Player1Name TEXT,Player2Id INTEGER,Player2Name TEXT,Player1Number INTEGER,Player2Number INTEGER,Start INTEGER NOT NULL DEFAULT 0,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOT NULL DEFAULT 0);`,
	}
	//open pseudo database for function
	var err error
	a.DB, err = sql.Open("ramsql", "TestCreateGameHandler")
	if err != nil {
		t.Fatalf("Error creating mock sql : %s\n", err)
	}
	defer a.DB.Close()

	// Exec every line of batch and create database
	for _, b := range batches {
		_, err = a.DB.Exec(b)
		if err != nil {
			t.Fatalf("Error exec query in query: %s\n Error:%s", b, err)
		}
	}
/////////////////////////////////// MOCKING ///////////////////////////////////////////
	tests := []struct {
		input  string
		output string
	}{

		{input: `{"name":"anil"}`, output: `{"id":0,"session":{"id":0,"date":"`+now+`","turn":0,"player1":{"id":0,"name":"anil"},"player2":{"id":0,"name":""},"player1number":0,"player2number":0,"start":0,"end":0,"winner":0},"clue":{"positive":0,"negative":0},"predictor":0,"prediction":0,"action":"Created"}`},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/CreateGame",  strings.NewReader(test.input))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.CreateGameHandler)

		handler.ServeHTTP(rr, req)

		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != test.output {
			t.Errorf("\nResponse is: %v .\nExpected :   %v", string(body), test.output)
		}
	}
}


func TestJoinGameHandler(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER DEFAULT 0,Negative INTEGER DEFAULT 0,Predictor INTEGER,Prediction INTEGER,Action TEXT);`,
		`INSERT INTO Moves (Id ,SessionId,Positive,Negative,Predictor,Prediction,Action) VALUES (1,1,0,0,0,0,"Created");`,
		`CREATE TABLE Users (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);`,
		`INSERT INTO Users (Id,Name) VALUES (1,"anonymous");`,
		`CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,Turn INTEGER NOT NULL,Player1Id INTEGER,Player1Name TEXT,Player2Id INTEGER,Player2Name TEXT,Player1Number INTEGER,Player2Number INTEGER,Start INTEGER NOT NULL DEFAULT 0,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOT NULL DEFAULT 0);`,
		`INSERT INTO Sessions (Id,Date,Turn,Player1Id,Player1Name,Player2Id,Player2Name,Player1Number,Player2Number,Start,End,Winner) VALUES (1,"`+now+`",0,1,"anonymous",0,"",0,0,0,0,0);`,
	}
	//open pseudo database for function
	var err error
	a.DB, err = sql.Open("ramsql", "TestJoinGameHandler")
	if err != nil {
		t.Fatalf("Error creating mock sql : %s\n", err)
	}
	defer a.DB.Close()

	// Exec every line of batch and create database
	for _, b := range batches {
		_, err = a.DB.Exec(b)
		if err != nil {
			t.Fatalf("Error exec query in query: %s\n Error:%s", b, err)
		}
	}
/////////////////////////////////// MOCKING ///////////////////////////////////////////
	tests := []struct {
		input  string
		output string
	}{

		{input: `{"user":{"name":"anil"},"session":{"id": 1}}`, output: `{"id":0,"session":{"id":1,"date":"`+now+`","turn":0,"player1":{"id":1,"name":"anonymous"},"player2":{"id":0,"name":"anil"},"player1number":0,"player2number":0,"start":0,"end":0,"winner":0},"clue":{"positive":0,"negative":0},"predictor":0,"prediction":0,"action":"Joined"}`},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/JoinGame",  strings.NewReader(test.input))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.JoinGameHandler)

		handler.ServeHTTP(rr, req)

		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != test.output {
			t.Errorf("\nResponse is: %v .\nExpected :   %v", string(body), test.output)
		}
	}
}


func TestGetReadyHandler(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER DEFAULT 0,Negative INTEGER DEFAULT 0,Predictor INTEGER,Prediction INTEGER,Action TEXT);`,
		`INSERT INTO Moves (Id ,SessionId,Positive,Negative,Predictor,Prediction,Action) VALUES (1,1,0,0,0,0,"Created");`,
		`INSERT INTO Moves (Id ,SessionId,Positive,Negative,Predictor,Prediction,Action) VALUES (2,1,0,0,0,0,"Joined");`,
		`CREATE TABLE Users (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);`,
		`INSERT INTO Users (Id,Name) VALUES (1,"anonymous");`,
		`INSERT INTO Users (Id,Name) VALUES (2,"test");`,
		`CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,Turn INTEGER NOT NULL,Player1Id INTEGER,Player1Name TEXT,Player2Id INTEGER,Player2Name TEXT,Player1Number INTEGER,Player2Number INTEGER,Start INTEGER NOT NULL DEFAULT 0,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOT NULL DEFAULT 0);`,
		`INSERT INTO Sessions (Id,Date,Turn,Player1Id,Player1Name,Player2Id,Player2Name,Player1Number,Player2Number,Start,End,Winner) VALUES (1,"`+now+`",0,1,"anonymous",2,"test",0,0,0,0,0);`,
	}
	//open pseudo database for function
	var err error
	a.DB, err = sql.Open("ramsql", "TestGetReadyHandler")
	if err != nil {
		t.Fatalf("Error creating mock sql : %s\n", err)
	}
	defer a.DB.Close()

	// Exec every line of batch and create database
	for _, b := range batches {
		_, err = a.DB.Exec(b)
		if err != nil {
			t.Fatalf("Error exec query in query: %s\n Error:%s", b, err)
		}
	}
/////////////////////////////////// MOCKING ///////////////////////////////////////////
	tests := []struct {
		input  string
		output string
	}{

		{input: `{"number":1111,"user":1,"session":1}`, output: `{"id":0,"session":{"id":1,"date":"`+now+`","turn":0,"player1":{"id":1,"name":"anonymous"},"player2":{"id":2,"name":"test"},"player1number":1111,"player2number":0,"start":1,"end":0,"winner":0},"clue":{"positive":0,"negative":0},"predictor":0,"prediction":0,"action":"Ready1"}`},
		{input: `{"number":2222,"user":2,"session":1}`, output: `{"id":0,"session":{"id":1,"date":"`+now+`","turn":0,"player1":{"id":1,"name":"anonymous"},"player2":{"id":2,"name":"test"},"player1number":0,"player2number":2222,"start":3,"end":0,"winner":0},"clue":{"positive":0,"negative":0},"predictor":1,"prediction":0,"action":"Started"}`},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/backend/GetReady",  strings.NewReader(test.input))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(a.GetReadyHandler)

		handler.ServeHTTP(rr, req)

		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != test.output {
			t.Errorf("\nResponse is: %v .\nExpected :  %v", string(body), test.output)
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
