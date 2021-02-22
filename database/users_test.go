//go test ./database/... -v --cover
package database

import (
	"testing"

	//import models package
	"github.com/anilkusc/BullsAndCows/models"
	"github.com/anilkusc/BullsAndCows/test"

	//import mocking 3. party library
	_ "github.com/proullon/ramsql/driver"
)

// You need to create a user object for using user methods.
var u User

//Test for ReadUser function.
func TestReadUser(t *testing.T) {
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
		{id: 2, result: models.User{Id: 2, Name: "testuser"}, err: nil},
	}
	// Create Database for this function.It defined in test/test.go file
	db := test.CreateDatabase(t, "TestGetUser")

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
