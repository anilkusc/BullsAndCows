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
		`INSERT INTO Users (Name) VALUES ('test');`,
		`CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,Start INTEGER NOT NULL DEFAULT 0,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOT NULL DEFAULT 0);`,
		`INSERT INTO Sessions (Date) VALUES ('` + now + `');`,
		//It is not support foreign key
		`CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER DEFAULT 0,Negative INTEGER DEFAULT 0,Turn INTEGER NOT NULL,Player1Id INTEGER,Player1Name TEXT,Player2Id INTEGER,Player2Name TEXT,Player1Number INTEGER,Player2Number INTEGER,Predictor INTEGER,Prediction INTEGER,Action TEXT);`,
		`INSERT INTO Moves (SessionId,Turn,Player1Id,Player1Name,Player2Id,Player2Name,Player1Number,Player2Number,Predictor,Prediction,Action) VALUES (1,2,10,'Player1',11,'Player2','0000','0000',1,1111,'Created');`,
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
