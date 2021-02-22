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

	var user models.Session
	return user, nil

}
func (s *Session) ReadSession(db *sql.DB, session models.Session) (models.Session, error) {

	var user models.Session
	return user, nil

}
func (s *Session) UpdateSession(db *sql.DB, session models.Session) (models.Session, error) {

	var user models.Session
	return user, nil

}
func (s *Session) DeleteSession(db *sql.DB, session models.Session) (models.Session, error) {

	var user models.Session
	return user, nil

}
