package database

import (
	"database/sql"
	"strconv"

	models "github.com/anilkusc/BullsAndCows/models"
	_ "github.com/mattn/go-sqlite3"
)

type Move struct {
	*models.Move
}

func (m *Move) CreateMove(db *sql.DB, move models.Move) (models.Move, error) {
	statement, err := db.Prepare("INSERT INTO Moves (SessionId,Positive,Negative,Prediction,Action,Predictor,Start,End,Winner,Turn) VALUES(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return move, err
	}
	res, _ := statement.Exec(move.Session.Id, move.Clue.Positive, move.Clue.Negative, move.Prediction, move.Action, move.Session.Predictor, move.Session.Start, move.Session.End, move.Session.Winner, move.Session.Turn)
	statement.Close()
	id, _ := res.LastInsertId()
	move.Id = int(id)
	return move, nil

}
func (m *Move) ReadMove(db *sql.DB, id int) (models.Move, error) {
	var move models.Move

	query := "SELECT * FROM Moves where Id=" + strconv.Itoa(id)
	rows, err := db.Query(query)
	if err != nil {
		return move, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&move.Id, &move.Session.Id, &move.Clue.Positive, &move.Clue.Negative, &move.Prediction, &move.Action, &move.Session.Predictor, &move.Session.Start, &move.Session.End, &move.Session.Winner, &move.Session.Turn)
		if err != nil {
			return move, err
		}
	}
	if err = rows.Err(); err != nil {
		return move, err
	}
	return move, nil

}
func (m *Move) UpdateMove(db *sql.DB, move models.Move) (models.Move, error) {

	statement, err := db.Prepare("UPDATE Moves SET SessionId=?,Positive=?,Negative=?,Prediction=?,Action=?,Predictor=?,Start=?,End=?,Winner=?,Turn=? where Id=?")
	if err != nil {
		return move, err
	}
	statement.Exec(move.Session.Id, move.Clue.Positive, move.Clue.Negative, move.Prediction, move.Action, move.Session.Predictor, move.Session.Start, move.Session.End, move.Session.Winner, move.Session.Turn, move.Id)
	statement.Close()

	return move, nil

}
func (m *Move) DeleteMove(db *sql.DB, id int) (models.Move, error) {
	var move models.Move
	move, err := m.ReadMove(db, id)
	if err != nil {
		return move, err
	}

	statement, err := db.Prepare("DELETE FROM Moves where Id=?")
	if err != nil {
		return move, err
	}
	statement.Exec(move.Id)
	statement.Close()
	return move, nil

}
func (m *Move) ListMoves(db *sql.DB, sessionId int) ([]models.Move, error) {
	var moves []models.Move
	query := "SELECT * FROM Moves where SessionId=" + strconv.Itoa(sessionId)
	rows, err := db.Query(query)
	if err != nil {
		return moves, err
	}
	defer rows.Close()

	for rows.Next() {
		var move models.Move
		err := rows.Scan(&move.Id, &move.Session.Id, &move.Clue.Positive, &move.Clue.Negative, &move.Prediction, &move.Action, &move.Session.Predictor, &move.Session.Start, &move.Session.End, &move.Session.Winner, &move.Session.Turn)
		if err != nil {
			return moves, err
		}
		moves = append(moves, move)
	}
	if err = rows.Err(); err != nil {
		return moves, err
	}
	return moves, nil

}
