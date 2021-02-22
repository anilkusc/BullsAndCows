package database

import (
	"database/sql"

	models "github.com/anilkusc/BullsAndCows/models"
	_ "github.com/mattn/go-sqlite3"
)

type Session struct {
	*models.Session
}

func (s *Session) CreateSession(db *sql.DB, session models.Session) (models.Session, error) {

	return session, nil

}
func (s *Session) ReadSession(db *sql.DB, id int) (models.Session, error) {

	var session models.Session
	return session, nil

}
func (s *Session) UpdateSession(db *sql.DB, session models.Session) (models.Session, error) {

	return session, nil

}
func (s *Session) DeleteSession(db *sql.DB, id int) (models.Session, error) {

	var session models.Session
	return session, nil

}
