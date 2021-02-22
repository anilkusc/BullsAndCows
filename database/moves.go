package database

import (
	"database/sql"

	models "github.com/anilkusc/BullsAndCows/models"
	_ "github.com/mattn/go-sqlite3"
)

type Move struct {
	*models.Move
}

func (m *Move) CreateMove(db *sql.DB, move models.Move) (models.Move, error) {

	return move, nil

}
func (m *Move) ReadMove(db *sql.DB, id int) (models.Move, error) {
	var move models.Move
	return move, nil

}
func (m *Move) UpdateMove(db *sql.DB, move models.Move) (models.Move, error) {
	return move, nil

}
func (m *Move) DeleteMove(db *sql.DB, id int) (models.Move, error) {
	var move models.Move
	return move, nil

}
