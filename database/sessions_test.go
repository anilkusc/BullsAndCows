// for running =>  //go test ./database/... -v --cover
package database

import (
	"testing"
	"time"

	"github.com/anilkusc/BullsAndCows/models"
	"github.com/anilkusc/BullsAndCows/test"

	_ "github.com/proullon/ramsql/driver"
)

var now = time.Now().Format("02-Jan-2006")
var s Session

func TestCreateSession(t *testing.T) {

	tests := []struct {
		session models.Session
		result  models.Session
		err     error
	}{

		{session: models.Session{Date: now, Start: 0, End: 0, Winner: 0}, result: models.Session{Id: 1, Date: now, Start: 0, End: 0, Winner: 0}, err: nil},
	}
	db := test.CreateDatabase(t, "TestCreateSession")

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
	db := test.CreateDatabase(t, "TestReadSession")

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
		session models.Session
		result  models.Session
		err     error
	}{
		{session: models.Session{Id: 1, Date: now, Start: 0, End: 0, Winner: 0}, result: models.Session{Id: 1, Date: now, Start: 0, End: 0, Winner: 0}, err: nil},
	}
	db := test.CreateDatabase(t, "TestUpdateSession")

	defer db.Close()

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
	tests := []struct {
		id     int
		result models.Session
		err    error
	}{
		{id: 1, result: models.Session{Id: 1, Date: now, Start: 0, End: 0, Winner: 0}, err: nil},
	}
	db := test.CreateDatabase(t, "TestDeleteSession")

	defer db.Close()

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
