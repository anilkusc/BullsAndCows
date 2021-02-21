// for running =>  go test ./tests/...
// name should be database_test because we are trying to test database package
package database_test

import (
	"database/sql"
	"testing"

	//import database package
	database "github.com/anilkusc/BullsAndCows/database"
	"github.com/anilkusc/BullsAndCows/models"

	//import models package

	//import mocking 3. party library
	_ "github.com/proullon/ramsql/driver"
)

type User struct {
	*database.User
}

// You need to create a user object for using user methods.
var u User

// Create a database for all test functions.It takes a t and functionName parameterfrom original testing function(from the function it is in).
func CreateDatabase(t *testing.T, functionName string) *sql.DB {
	// Create pseudo database with queries.
	var batches = []string{
		`CREATE TABLE Users (Id INT PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);`,
		`INSERT INTO Users (Id,Name) VALUES (1, 'anonymous');`,
		`INSERT INTO Users (Id,Name) VALUES (2, 'testuser');`,
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

//Test for ReadUser function.
func TestReadUser(t *testing.T) {
	// Specify test variables and expected results.
	tests := []struct {
		id     int
		result models.User
		err    error
	}{
		// When give to first parameter(id) 1 , We expect result :1 error nil
		{id: 1, result: models.User{Id: 1, Name: "anonymous"}, err: nil},
		// When give to first parameter(id) 1 , We expect result :1 error nil
		{id: 2, result: models.User{Id: 2, Name: "testuser"}, err: nil},
	}
	// Create Database for this function.
	db := CreateDatabase(t, "TestGetUser")

	defer db.Close()

	// test all of the variables.
	for _, test := range tests {
		//get result after test.
		s, err := u.ReadUser(db, test.id)
		// if expected error type nil we need to compare with actual error different way.
		if test.err == nil {
			// If test fails give error.It checks expected result and expected error
			if err != test.err || s != test.result {
				// Compare expected error and actual error
				t.Errorf("Error is: %v . Expected: %v", test.err, err)
				// Compare expected result and actual result
				t.Errorf("Result is: %v . Expected: %v", test.result, s)
			}
			// if expected error type is not nil we need to compare with actual error different way.
		} else {
			if err.Error() != test.err.Error() || s != test.result {
				// Compare expected error and actual error
				t.Errorf("Error is: %v . Expected: %v", test.err, err)
				// Compare expected result and actual result
				t.Errorf("Result is: %v . Expected: %v", test.result, s)
			}
		}
	}
}
