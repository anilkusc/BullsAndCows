//go test ./database/... -v --cover
package database

import (
	"reflect"
	"testing"

	"github.com/anilkusc/BullsAndCows/models"
	"github.com/anilkusc/BullsAndCows/test"

	_ "github.com/proullon/ramsql/driver"
)

var m Move

func TestCreateMove(t *testing.T) {

	tests := []struct {
		move   models.Move
		result models.Move
		err    error
	}{
		{move: models.Move{Session: models.Session{Id: 1}, Turn: 0, Action: "Created"}, result: models.Move{Id: 0, Session: models.Session{Id: 1}, Clue: models.Clue{Positive: 0, Negative: 0}, Turn: 0, Player1: models.User{Id: 0, Name: ""}, Player2: models.User{Id: 0, Name: ""}, Player1Number: 0, Player2Number: 0, Predictor: 0, Prediction: 0, Action: "Created"}, err: nil},
	}
	db := test.CreateDatabase(t, "TestCreateMove")

	defer db.Close()

	for _, test := range tests {
		s, err := m.CreateMove(db, test.move)
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

		{id: 1, result: models.Move{Id: 1, Session: models.Session{Id: 1}, Clue: models.Clue{Positive: 0, Negative: 0}, Turn: 2, Player1: models.User{Id: 10, Name: "Player1"}, Player2: models.User{Id: 11, Name: "Player2"}, Player1Number: 0, Player2Number: 0, Predictor: 1, Prediction: 1111, Action: "Created"}, err: nil},
	}
	db := test.CreateDatabase(t, "TestReadMove")

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
		move   models.Move
		result models.Move
		err    error
	}{

		{move: models.Move{Id: 1, Session: models.Session{Id: 1}, Clue: models.Clue{Positive: 0, Negative: 0}, Turn: 2, Player1: models.User{Id: 10, Name: "Player1"}, Player2: models.User{Id: 11, Name: "Player2"}, Player1Number: 0, Player2Number: 0, Predictor: 1, Prediction: 2222, Action: "Updated"}, result: models.Move{Id: 1, Session: models.Session{Id: 1}, Clue: models.Clue{Positive: 0, Negative: 0}, Turn: 2, Player1: models.User{Id: 10, Name: "Player1"}, Player2: models.User{Id: 11, Name: "Player2"}, Player1Number: 0, Player2Number: 0, Predictor: 1, Prediction: 2222, Action: "Updated"}, err: nil},
	}
	db := test.CreateDatabase(t, "TestUpdateMove")

	defer db.Close()

	for _, test := range tests {
		s, err := m.UpdateMove(db, test.move)
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

		{id: 1, result: models.Move{Id: 1, Session: models.Session{Id: 1}, Clue: models.Clue{Positive: 0, Negative: 0}, Turn: 2, Player1: models.User{Id: 10, Name: "Player1"}, Player2: models.User{Id: 11, Name: "Player2"}, Player1Number: 0, Player2Number: 0, Predictor: 1, Prediction: 1111, Action: "Created"}, err: nil},
	}
	db := test.CreateDatabase(t, "TestDeleteMove")

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
func TestListMoves(t *testing.T) {

	tests := []struct {
		id     int
		result []models.Move
		err    error
	}{

		{id: 1, result: []models.Move{{Id: 1, Session: models.Session{Id: 1}, Clue: models.Clue{Positive: 0, Negative: 0}, Turn: 2, Player1: models.User{Id: 10, Name: "Player1"}, Player2: models.User{Id: 11, Name: "Player2"}, Player1Number: 0, Player2Number: 0, Predictor: 1, Prediction: 1111, Action: "Created"}}, err: nil},
	}
	db := test.CreateDatabase(t, "TestListMoves")

	defer db.Close()

	for _, test := range tests {
		s, err := m.ListMoves(db, test.id)
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
