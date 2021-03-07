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
	statement, err := db.Prepare("INSERT INTO Sessions VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return move, err
	}
	statement.Exec(move.Session.Id, move.Clue.Positive, move.Clue.Negative, move.Turn, move.Player1.Id, move.Player1.Name, move.Player2.Id, move.Player2.Name, move.Player1Number, move.Player2Number, move.Predictor, move.Prediction, move.Action)
	statement.Close()
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
		err := rows.Scan(&move.Id, &move.Session.Id, &move.Clue.Positive, &move.Clue.Negative, &move.Turn, &move.Player1.Id, &move.Player1.Name, &move.Player2.Id, &move.Player2.Name, &move.Player1Number, &move.Player2Number, &move.Predictor, &move.Prediction, &move.Action)
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
	statement, err := db.Prepare("UPDATE Moves SET SessionId=?,Positive=?,Negative=?,Turn=?,Player1Id=?,Player1Name=?,Player2Id=?,Player2Name=?,Player1Number=?,Player2Number=?,Predictor=?,Prediction=?,Action=?, where Id=?")
	if err != nil {
		return move, err
	}
	statement.Exec(move.Session.Id, move.Clue.Positive, move.Clue.Negative, move.Turn, move.Player1.Id, move.Player1.Name, move.Player2.Id, move.Player2.Name, move.Player1Number, move.Player2Number, move.Predictor, move.Prediction, move.Action, move.Id)
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
		err := rows.Scan(&move.Id, &move.Session.Id, &move.Clue.Positive, &move.Clue.Negative, &move.Turn, &move.Player1.Id, &move.Player1.Name, &move.Player2.Id, &move.Player2.Name, &move.Player1Number, &move.Player2Number, &move.Predictor, &move.Prediction, &move.Action)
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
