package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/proullon/ramsql/driver"
)

var a App

func TestReadUserHandler(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{

		{input: "test", output: "test"},
	}
	//db := CreateDatabase(t, "TestReadUserHandler")
	//defer db.Close()

	//request := httptest.NewRequest("POST", "/backend/ReadUser", nil)
	//responseRecorder := httptest.NewRecorder()
	//ReadUser{test.input}.ServeHTTP(responseRecorder, request)

	for _, test := range tests {
		//s, err := s.CreateSession(db, test.session)

		req, _ := http.NewRequest("POST", "/backend/ReadUserHandler", nil)
		res := httptest.NewRecorder()
		a.Router.ServeHTTP(res, req)

		body, _ := ioutil.ReadAll(res.Body)
		if string(body) != test.output {
			t.Errorf("Response is: %v . Expected: %v", res, test.output)
		}
	}
}
