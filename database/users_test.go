//go test ./database/... -v --cover
package database

import (
	"database/sql"
	"reflect"
	"testing"

	//import models package
	"github.com/anilkusc/BullsAndCows/models"

	//import mocking 3. party library
	_ "github.com/proullon/ramsql/driver"
)

// You need to create a user object for using user methods.
var u User

func TestCreateUser(t *testing.T) {
//It is better to mocking again for every function for detailed testing.
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Users (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);`,
	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestCreateUser")
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
		user   models.User
		result models.User
		err    error
	}{
		{user: models.User{Name: "myuser"}, result: models.User{Id: 0, Name: "myuser"}, err: nil},
	}

	for _, test := range tests {
		s, err := u.CreateUser(db, test.user)
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

//Test for ReadUser function.
func TestReadUser(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Users (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);`,
		`INSERT INTO Users (Id,Name) VALUES (1,'anonymous');`,
	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestReadUser")
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
/////////////////////////////////// MOCKING ////////////////////////////////////////////

	// Specify test variables and expected results.
	tests := []struct {
		id int
		// we need to use models.User for passing to object.This is different with "database.User".
		result models.User
		err    error
	}{
		// When give to first parameter(id) 1 , We expect result :1 error nil
		{id: 1, result: models.User{Id: 1, Name: "anonymous"}, err: nil},
		// When give to first parameter(id) 1 , We expect result :1 error nil
		//{id: 2, result: models.User{Id: 2, Name: "test"}, err: nil},
	}

	// test all of the variables.
	for _, test := range tests {
		//get result after test.
		s, err := u.ReadUser(db, test.id)
		// if expected error type nil we need to compare with actual error different way.
		if test.err == nil {
			// If test fails give error.It checks expected result and expected error
			if err != test.err || s != test.result {
				// Compare expected error and actual error
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				// Compare expected result and actual result
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
			// if expected error type is not nil we need to compare with actual error different way.
		} else {
			if err.Error() != test.err.Error() || s != test.result {
				// Compare expected error and actual error
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				// Compare expected result and actual result
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		}
	}
}

func TestUpdateUser(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Users (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);`,
		`INSERT INTO Users (Id,Name) VALUES (1,'anonymous');`,
	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestUpdateUser")
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
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	tests := []struct {
		user   models.User
		result models.User
		err    error
	}{
		{user: models.User{Id: 1, Name: "John"}, result: models.User{Id: 1, Name: "John"}, err: nil},
	}

	for _, test := range tests {
		s, err := u.UpdateUser(db, test.user)
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

func TestDeleteUser(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Users (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);`,
		`INSERT INTO Users (Id,Name) VALUES (1,'anonymous');`,
	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestDeleteUser")
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
/////////////////////////////////// MOCKING ////////////////////////////////////////////	
	tests := []struct {
		id     int
		result models.User
		err    error
	}{
		{id: 1, result: models.User{Id: 1, Name: "anonymous"}, err: nil},
	}

	for _, test := range tests {
		s, err := u.DeleteUser(db, test.id)
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

func TestListUsers(t *testing.T) {
/////////////////////////////////// MOCKING ////////////////////////////////////////////
	var batches = []string{
		`CREATE TABLE Users (Id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, Name TEXT NOT NULL UNIQUE);`,
		`INSERT INTO Users (Id,Name) VALUES (1,'anonymous');`,
		`INSERT INTO Users (Id,Name) VALUES (2,'test');`,
	}
	//open pseudo database for function
	db, err := sql.Open("ramsql", "TestListUsers")
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
/////////////////////////////////// MOCKING ////////////////////////////////////////////	
	tests := []struct {
		result []models.User
		err    error
	}{
		{result: []models.User{{Id: 1, Name: "anonymous"}, {Id: 2, Name: "test"}}, err: nil},
	}

	for _, test := range tests {
		s, err := u.ListUsers(db)
		if test.err == nil {
			//you need to user reflect for compare 2 object array
			if err != test.err || reflect.DeepEqual(s, test.result) != true {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		} else {
			if err.Error() != test.err.Error() || reflect.DeepEqual(s, test.result) {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
				t.Errorf("Result is: %v . Expected: %v", s, test.result)
			}
		}
	}
}
