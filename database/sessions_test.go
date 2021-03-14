// for running =>  //go test ./database/... -v --cover
package database

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/anilkusc/BullsAndCows/models"

	_ "github.com/proullon/ramsql/driver"
)

var now = time.Now().Format("02-Jan-2006")
var s Session

func TestCreateSession(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,Turn INTEGER NOT NULL,Player1Id INTEGER,Player1Name TEXT,Player2Id INTEGER,Player2Name TEXT,Player1Number INTEGER,Player2Number INTEGER,Start INTEGER NOT NULL DEFAULT 0,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOT NULL DEFAULT 0);`,
	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestCreateSession")
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
		session models.Session
		result  models.Session
		err     error
	}{

		{session: models.Session{Date: now,Turn:0,Player1:models.User{Id:1,Name:"Anonymous1"},Player2:models.User{Id:0,Name:""},Player1Number:0,Player2Number:0,Start:0,End:0,Winner:0}, result: models.Session{Date: now,Turn:0,Player1:models.User{Id:1,Name:"Anonymous1"},Player2:models.User{Id:0,Name:""},Player1Number:0,Player2Number:0,Start:0,End:0,Winner:0}, err: nil},
	}

	for _, test := range tests {
		s, err := s.CreateSession(db, test.session)
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
func TestReadSession(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,Turn INTEGER NOT NULL,Player1Id INTEGER,Player1Name TEXT,Player2Id INTEGER,Player2Name TEXT,Player1Number INTEGER,Player2Number INTEGER,Start INTEGER NOT NULL DEFAULT 0,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOT NULL DEFAULT 0);`,
		`INSERT INTO Sessions (Id,Date,Turn,Player1Id,Player1Name,Player2Id,Player2Name,Player1Number,Player2Number,Start,End,Winner) VALUES (1,'`+now+`',5,1,'anonymous',2,'test',1111,2222,1,0,0);`,

	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestReadSession")
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
		result models.Session
		err    error
	}{
		{id: 1, result: models.Session{Id: 1, Date: now,Turn: 5,Player1: models.User{Id: 1, Name: "anonymous"},Player2: models.User{Id: 2, Name: "test"},Player1Number: 1111,Player2Number: 2222, Start: 1, End: 0, Winner: 0}, err: nil},
	}

	for _, test := range tests {
		s, err := s.ReadSession(db, test.id)
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

func TestUpdateSession(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,Turn INTEGER NOT NULL,Player1Id INTEGER,Player1Name TEXT,Player2Id INTEGER,Player2Name TEXT,Player1Number INTEGER,Player2Number INTEGER,Start INTEGER NOT NULL DEFAULT 0,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOT NULL DEFAULT 0);`,
		`INSERT INTO Sessions (Id,Date,Turn,Player1Id,Player1Name,Player2Id,Player2Name,Player1Number,Player2Number,Start,End,Winner) VALUES (1,'`+now+`',5,1,'anonymous',2,'test',1111,2222,1,0,0);`,
	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestUpdateSession")
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
		session models.Session
		result  models.Session
		err     error
	}{
		{session: models.Session{Id: 1, Date: now,Turn: 5,Player1: models.User{Id: 1, Name: "myuser"},Player2: models.User{Id: 2, Name: "test"},Player1Number: 1111,Player2Number: 2222, Start: 1, End: 0, Winner: 0}, result: models.Session{Id: 1, Date: now,Turn: 5,Player1: models.User{Id: 1, Name: "myuser"},Player2: models.User{Id: 2, Name: "test"},Player1Number: 1111,Player2Number: 2222, Start: 1, End: 0, Winner: 0}, err: nil},
	}

	for _, test := range tests {
		s, err := s.UpdateSession(db, test.session)
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
func TestDeleteSession(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,Turn INTEGER NOT NULL,Player1Id INTEGER,Player1Name TEXT,Player2Id INTEGER,Player2Name TEXT,Player1Number INTEGER,Player2Number INTEGER,Start INTEGER NOT NULL DEFAULT 0,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOT NULL DEFAULT 0);`,
		`INSERT INTO Sessions (Id,Date,Turn,Player1Id,Player1Name,Player2Id,Player2Name,Player1Number,Player2Number,Start,End,Winner) VALUES (1,'`+now+`',5,1,'anonymous',2,'test',1111,2222,1,0,0);`,

	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestDeleteSession")
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
		result models.Session
		err    error
	}{
		{id: 1, result: models.Session{Id: 1, Date: now,Turn: 5,Player1: models.User{Id: 1, Name: "anonymous"},Player2: models.User{Id: 2, Name: "test"},Player1Number: 1111,Player2Number: 2222, Start: 1, End: 0, Winner: 0}, err: nil},
	}

	for _, test := range tests {
		s, err := s.DeleteSession(db, test.id)
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
func TestListSessions(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,Turn INTEGER NOT NULL,Player1Id INTEGER,Player1Name TEXT,Player2Id INTEGER,Player2Name TEXT,Player1Number INTEGER,Player2Number INTEGER,Start INTEGER NOT NULL DEFAULT 0,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOT NULL DEFAULT 0);`,
		`INSERT INTO Sessions (Id,Date,Turn,Player1Id,Player1Name,Player2Id,Player2Name,Player1Number,Player2Number,Start,End,Winner) VALUES (1,'`+now+`',5,1,'anonymous',2,'test',1111,2222,1,0,0);`,
		`INSERT INTO Sessions (Id,Date,Turn,Player1Id,Player1Name,Player2Id,Player2Name,Player1Number,Player2Number,Start,End,Winner) VALUES (2,'`+now+`',3,1,'anonymous',2,'test',3333,4444,1,0,0);`,

	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestListSessions")
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
		result []models.Session
		err    error
	}{
		{result: []models.Session{{Id: 1, Date: now,Turn: 5,Player1: models.User{Id: 1, Name: "anonymous"},Player2: models.User{Id: 2, Name: "test"},Player1Number: 1111,Player2Number: 2222, Start: 1, End: 0, Winner: 0},{Id: 2, Date: now,Turn: 3,Player1: models.User{Id: 1, Name: "anonymous"},Player2: models.User{Id: 2, Name: "test"},Player1Number: 3333,Player2Number: 4444, Start: 1, End: 0, Winner: 0}}, err: nil},
	}

	for _, test := range tests {
		s, err := s.ListSessions(db)
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
