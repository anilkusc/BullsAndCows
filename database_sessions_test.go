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

func TestCreateSession(t *testing.T) {
	tests := []struct {
		session models.Session
		result  models.Session
		err     error
	}{
		{session: models.Session{Date: time.Now().String(), Start: 0, End: 0, Winner: 0}, result: models.Session{Id: 1, Date: time.Now().String(), Start: 0, End: 0, Winner: 0}, err: nil},
	}
	db := CreateDatabase(t, "TestCreateSession")

	defer db.Close()

	for _, test := range tests {
		s, err := s.CreateSession(db, test.session)
		if test.err == nil {
			if err != test.err || s != test.result {
				t.Errorf("Error is: %v . Expected: %v", test.err, err)
				t.Errorf("Result is: %v . Expected: %v", test.result, s)
			}
		} else {
			if err.Error() != test.err.Error() || s != test.result {
				t.Errorf("Error is: %v . Expected: %v", test.err, err)
				t.Errorf("Result is: %v . Expected: %v", test.result, s)
			}
		}
	}
}
