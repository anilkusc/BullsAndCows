package database_test

import (
	"testing"

	database "github.com/anilkusc/BullsAndCows/database"

	"github.com/anilkusc/BullsAndCows/models"

	_ "github.com/proullon/ramsql/driver"
)

type Move struct {
	*database.Move
}

var m Move

func TestCreateMove(t *testing.T) {

	tests := []struct {
		session models.Move
		result  models.Move
		err     error
	}{

		{session: models.Move{Session: models.Session{Id: 1}, Turn: 0, Action: "Created"}, result: models.Move{Session: models.Session{Id: 1}, Turn: 0, Action: "Created"}, err: nil},
	}
	db := CreateDatabase(t, "TestCreateMove")

	defer db.Close()

	for _, test := range tests {
		s, err := m.CreateMove(db, test.session)
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

	tests := []struct {
		id     int
		result models.Move
		err    error
	}{

		{id: 1, result: models.Move{Session: models.Session{Id: 1}, Turn: 0, Action: "Created"}, err: nil},
	}
	db := CreateDatabase(t, "TestReadMove")

	defer db.Close()

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

	tests := []struct {
		session models.Move
		result  models.Move
		err     error
	}{

		{session: models.Move{Session: models.Session{Id: 1}, Turn: 0, Action: "Updated"}, result: models.Move{Session: models.Session{Id: 1}, Turn: 0, Action: "Created"}, err: nil},
	}
	db := CreateDatabase(t, "TestUpdateMove")

	defer db.Close()

	for _, test := range tests {
		s, err := m.UpdateMove(db, test.session)
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

	tests := []struct {
		id     int
		result models.Move
		err    error
	}{

		{id: 1, result: models.Move{Session: models.Session{Id: 1}, Turn: 0, Action: "Created"}, err: nil},
	}
	db := CreateDatabase(t, "TestDeleteMove")

	defer db.Close()

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
