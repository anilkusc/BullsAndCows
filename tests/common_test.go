package database_test

import (
	"database/sql"
	"testing"

	_ "github.com/proullon/ramsql/driver"
)

// Create a database for all test functions.It takes a t and functionName parameterfrom original testing function(from the function it is in).

func CreateDatabase(t *testing.T, functionName string) *sql.DB {
	// Create pseudo database with queries.
	var batches = []string{
		`CREATE TABLE Users (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);`,
		`INSERT INTO Users (Name) VALUES ('anonymous');`,
		`INSERT INTO Users (Name) VALUES ('testuser');`,
		`CREATE TABLE Sessions (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Date TEXT NOT NULL,End INTEGER NOT NULL DEFAULT 0,Winner INTEGER NOY NULL DEFAULT 0);`,
		`INSERT INTO Sessions (Date,Winner) VALUES ('Now');`,
		`CREATE TABLE Moves (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,SessionId INTEGER NOT NULL,Positive INTEGER DEFAULT -1,Negative INTEGER DEFAULT -1,Turn INTEGER NOT NULL,Player1 TEXT,Player2 TEXT,Player1Number INTEGER,Player2Number INTEGER,Predictor INTEGER,Prediction INTEGER,Action TEXT,FOREIGN KEY (SessionId) REFERENCES Sessions (Id) ON DELETE CASCADE);`,
		`INSERT INTO Moves (SessionId,Turn,Action) VALUES (1,0,'Created');`,
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
			t.Fatalf("Error exec query in mock query: %s\n", err)
		}
	}
	return db
}
