//go test ./database/... -v --cover
package test

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/proullon/ramsql/driver"
)

var now = time.Now().Format("02-Jan-2006")

// Create a database for all test functions.It takes a t and functionName parameterfrom original testing function(from the function it is in).

func CreateDatabase(t *testing.T, functionName string) *sql.DB {
	// Create pseudo database with queries.
	var batches = []string{
		`CREATE TABLE Users (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);`,
		`INSERT INTO Users (Name) VALUES ('anonymous');`,
		`INSERT INTO Users (Name) VALUES ('testuser');`,
		`CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOT NULL DEFAULT 0);`,
		//`INSERT INTO Sessions (Date,Winner) VALUES ('Now',0);`,
		//It is not support foreign key
		`CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER ,Negative INTEGER ,Turn INTEGER NOT NULL,Player1 TEXT,Player2 TEXT,Player1Number INTEGER,Player2Number INTEGER,Predictor INTEGER,Prediction INTEGER,Action TEXT);`,
		//`INSERT INTO Moves (SessionId,Turn,Action) VALUES (1,0,'Created');`,
	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", functionName)
	if err != nil {
		t.Fatalf("Error creating mock sql : %s\n", err)
	}
	// Exec every line of batch and create database
	for _, b := range batches {
		_, err = db.Exec(b)
		if err != nil {
			t.Fatalf("Error exec query in query: %s\n Error:%s", b, err)
		}
	}
	return db
}
