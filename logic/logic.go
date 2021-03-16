package logic

import (
	models "github.com/anilkusc/BullsAndCows/models"
)


func CalculateClue(prediction int, actualNumber int) (models.Clue,error) {
	var clue models.Clue
	clue.Positive = 2
	clue.Negative = 2
	return clue,nil
}
func CalculateWinner() int {
	return 0
}
