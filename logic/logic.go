package logic

import (
	models "github.com/anilkusc/BullsAndCows/models"
)

type Clue struct {
	Clue *models.Clue
}

func CalculateClue(prediction int, actualNumber int) Clue {
	var clue Clue
	return clue
}
func CalculateWinner() int {
	return 0
}
