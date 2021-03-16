//go test ./database/... -v --cover
package database

import (
	"database/sql"
	"reflect"
	"testing"

	models "github.com/anilkusc/BullsAndCows/models"

	_ "github.com/proullon/ramsql/driver"
)

var m Move

func TestCreateMove(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
var batches = []string{
	`CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER DEFAULT 0,Negative INTEGER DEFAULT 0,Prediction INTEGER,Action TEXT);`,
}
//open pseudo database for function
db, err := sql.Open("ramsql", "TestCreateMove")
if err != nil {
	t.Fatalf("Error creating mock sql : %s\n", err)
}
defer db.Close()

// Exec every line of batch and create database
for _, b := range batches {
	_, err = db.Exec(b)
	if err != nil {
		t.Fatalf("Error exec query in query: %s\n Error:%s", b, err)
	}
}
/////////////////////////////////// MOCKING ///////////////////////////////////////////

	tests := []struct {
		move   models.Move
		result models.Move
		err    error
	}{
		{ move: models.Move{Id: 10,Session: models.Session{Id:1},Clue: models.Clue{Positive: 2,Negative:2},Prediction: 2222,Action: "Predicted"}, result: models.Move{Id: 10,Session: models.Session{Id:1},Clue: models.Clue{Positive: 2,Negative:2},Prediction: 2222,Action: "Predicted"}, err: nil},
	}

	for _, test := range tests {
		s, err := m.CreateMove(db, test.move)
		if test.err == nil {
			if err != test.err || s != test.result {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		} else {
			if err.Error() != test.err.Error() || s != test.result {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		}
	}
}
func TestReadMove(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER DEFAULT 0,Negative INTEGER DEFAULT 0,Prediction INTEGER,Action TEXT);`,
		`INSERT INTO Moves (Id,SessionId,Positive,Negative,Prediction,Action) VALUES (1,1,3,4,1111,"Predicted");`,
	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestReadMove")
	if err != nil {
		t.Fatalf("Error creating mock sql : %s\n", err)
	}
	defer db.Close()

	// Exec every line of batch and create database
	for _, b := range batches {
		_, err = db.Exec(b)
		if err != nil {
			t.Fatalf("Error exec query in query: %s\n Error:%s", b, err)
		}
	}
/////////////////////////////////// MOCKING ///////////////////////////////////////////

	tests := []struct {
		id     int
		result models.Move
		err    error
	}{

		{id: 1, result: models.Move{Id: 1, Session: models.Session{Id: 1}, Clue: models.Clue{Positive: 3, Negative: 4}, Prediction: 1111, Action: "Predicted"}, err: nil},
	}

	for _, test := range tests {
		s, err := m.ReadMove(db, test.id)
		if test.err == nil {
			if err != test.err || s != test.result {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		} else {
			if err.Error() != test.err.Error() || s != test.result {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		}
	}
}

func TestUpdateMove(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER DEFAULT 0,Negative INTEGER DEFAULT 0,Prediction INTEGER,Action TEXT);`,
		`INSERT INTO Moves (Id,SessionId,Positive,Negative,Prediction,Action) VALUES (1,1,3,4,1111,"Predicted");`,
	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestUpdateMove")
	if err != nil {
		t.Fatalf("Error creating mock sql : %s\n", err)
	}
	defer db.Close()

	// Exec every line of batch and create database
	for _, b := range batches {
		_, err = db.Exec(b)
		if err != nil {
			t.Fatalf("Error exec query in query: %s\n Error:%s", b, err)
		}
	}
/////////////////////////////////// MOCKING ///////////////////////////////////////////
	tests := []struct {
		move   models.Move
		result models.Move
		err    error
	}{

		{move: models.Move{Id: 1, Session: models.Session{Id: 1}, Clue: models.Clue{Positive: 1, Negative: 1}, Prediction: 2222, Action: "Predicted"},result: models.Move{Id: 1, Session: models.Session{Id: 1}, Clue: models.Clue{Positive: 1, Negative: 1}, Prediction: 2222, Action: "Predicted"} ,err: nil},
	}

	for _, test := range tests {
		s, err := m.UpdateMove(db, test.move)
		if test.err == nil {
			if err != test.err || s != test.result {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		} else {
			if err.Error() != test.err.Error() || s != test.result {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		}
	}
}
func TestDeleteMove(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER DEFAULT 0,Negative INTEGER DEFAULT 0,Prediction INTEGER,Action TEXT);`,
		`INSERT INTO Moves (Id,SessionId,Positive,Negative,Prediction,Action) VALUES (1,1,3,4,1111,"Predicted");`,
	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestDeleteMove")
	if err != nil {
		t.Fatalf("Error creating mock sql : %s\n", err)
	}
	defer db.Close()

	// Exec every line of batch and create database
	for _, b := range batches {
		_, err = db.Exec(b)
		if err != nil {
			t.Fatalf("Error exec query in query: %s\n Error:%s", b, err)
		}
	}
/////////////////////////////////// MOCKING ///////////////////////////////////////////
	tests := []struct {
		id     int
		result models.Move
		err    error
	}{

		{id: 1, result: models.Move{Id: 1, Session: models.Session{Id: 1}, Clue: models.Clue{Positive: 3, Negative: 4},Prediction: 1111, Action: "Predicted"}, err: nil},
	}

	for _, test := range tests {
		s, err := m.DeleteMove(db, test.id)
		if test.err == nil {
			if err != test.err || s != test.result {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		} else {
			if err.Error() != test.err.Error() || s != test.result {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		}
	}
}
func TestListMoves(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER DEFAULT 0,Negative INTEGER DEFAULT 0,Prediction INTEGER,Action TEXT);`,
		`INSERT INTO Moves (Id,SessionId,Positive,Negative,Prediction,Action) VALUES (1,1,3,4,1111,"Predicted");`,
		`INSERT INTO Moves (Id,SessionId,Positive,Negative,Prediction,Action) VALUES (2,1,2,2,2222,"Predicted");`,
	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestListMoves")
	if err != nil {
		t.Fatalf("Error creating mock sql : %s\n", err)
	}
	defer db.Close()

	// Exec every line of batch and create database
	for _, b := range batches {
		_, err = db.Exec(b)
		if err != nil {
			t.Fatalf("Error exec query in query: %s\n Error:%s", b, err)
		}
	}
/////////////////////////////////// MOCKING ///////////////////////////////////////////
	tests := []struct {
		id     int
		result []models.Move
		err    error
	}{

		{id: 1, result: []models.Move{{Id: 1, Session: models.Session{Id: 1}, Clue: models.Clue{Positive: 3, Negative: 4}, Prediction: 1111, Action: "Predicted"},{Id: 2, Session: models.Session{Id: 1}, Clue: models.Clue{Positive: 2, Negative: 2}, Prediction: 2222, Action: "Predicted"}}, err: nil},
	}

	for _, test := range tests {
		s, err := m.ListMoves(db, test.id)
		if test.err == nil {
			if err != test.err || reflect.DeepEqual(s, test.result) != true {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		} else {
			if err.Error() != test.err.Error() || reflect.DeepEqual(s, test.result) != true {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		}
	}
}
