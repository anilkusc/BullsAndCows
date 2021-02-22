// for running =>  go test ./tests/...
// name should be database_test because we are trying to test database package
package database_test

import (
	"testing"
	"time"

	database "github.com/anilkusc/BullsAndCows/database"

	"github.com/anilkusc/BullsAndCows/models"

	_ "github.com/proullon/ramsql/driver"
)

type Session struct {
	*database.Session
}

var s Session
var now = time.Now().Format("02-Jan-2006")

func TestCreateSession(t *testing.T) {

	tests := []struct {
		session models.Session
		result  models.Session
		err     error
	}{

		{session: models.Session{Date: now, Start: 0, End: 0, Winner: 0}, result: models.Session{Id: 1, Date: now, Start: 0, End: 0, Winner: 0}, err: nil},
	}
	db := CreateDatabase(t, "TestCreateSession")

	defer db.Close()

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
	tests := []struct {
		id     int
		result models.Session
		err    error
	}{
		{id: 1, result: models.Session{Id: 1, Date: now, Start: 0, End: 0, Winner: 0}, err: nil},
	}
	db := CreateDatabase(t, "TestReadSession")

	defer db.Close()

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
	tests := []struct {
		id     int
		result models.Session
		err    error
	}{
		{id: 1, result: models.Session{Id: 1, Date: now, Start: 0, End: 0, Winner: 0}, err: nil},
	}
	db := CreateDatabase(t, "TestUpdateSession")

	defer db.Close()

	for _, test := range tests {
		s, err := s.UpdateSession(db, test.id)
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
	tests := []struct {
		id     int
		result models.Session
		err    error
	}{
		{id: 1, result: models.Session{Id: 1, Date: now, Start: 0, End: 0, Winner: 0}, err: nil},
	}
	db := CreateDatabase(t, "TestDeleteSession")

	defer db.Close()

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
