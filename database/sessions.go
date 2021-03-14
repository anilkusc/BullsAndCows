package database

import (
	"database/sql"
	"strconv"

	models "github.com/anilkusc/BullsAndCows/models"
	_ "github.com/mattn/go-sqlite3"
)

type Session struct {
	*models.Session
}

func (s *Session) CreateSession(db *sql.DB, session models.Session) (models.Session, error) {

	statement, err := db.Prepare("INSERT INTO Sessions (Date) VALUES(?)")
	if err != nil {
		return session, err
	}
	statement.Exec(session.Date)
	statement.Close()
	return session, nil

}
func (s *Session) ReadSession(db *sql.DB, id int) (models.Session, error) {

	var session models.Session

	query := "SELECT * FROM Sessions where Id='" + strconv.Itoa(id)+"'"
	rows, err := db.Query(query)
	if err != nil {
		return session, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&session.Id, &session.Date, &session.Start, &session.End, &session.Winner)
		if err != nil {
			return session, err
		}
	}
	if err = rows.Err(); err != nil {
		return session, err
	}
	return session, nil

}
func (s *Session) UpdateSession(db *sql.DB, session models.Session) (models.Session, error) {

	statement, err := db.Prepare("UPDATE Sessions SET Date=?,Start=?,End=?,Winner=? where Id=?")
	if err != nil {
		return session, err
	}
	statement.Exec(session.Date, session.Start, session.End, session.Winner, session.Id)
	statement.Close()

	return session, nil

}
func (s *Session) DeleteSession(db *sql.DB, id int) (models.Session, error) {

	var session models.Session
	session, err := s.ReadSession(db, id)
	if err != nil {
		return session, err
	}

	statement, err := db.Prepare("DELETE FROM Sessions where Id=?")
	if err != nil {
		return session, err
	}
	statement.Exec(session.Id)
	statement.Close()

	return session, nil

}
func (s *Session) ListSessions(db *sql.DB) ([]models.Session, error) {

	var sessions []models.Session

	query := "SELECT * FROM Sessions"
	rows, err := db.Query(query)
	if err != nil {
		return sessions, err
	}
	defer rows.Close()
	for rows.Next() {
		var session models.Session
		err := rows.Scan(&session.Id, &session.Date, &session.Start, &session.End, &session.Winner)
		if err != nil {
			return sessions, err
		}

		sessions = append(sessions, session)
	}
	if err = rows.Err(); err != nil {
		return sessions, err
	}
	return sessions, nil

}
