package logic

import (
	"strconv"

	models "github.com/anilkusc/BullsAndCows/models"
)

func CalculateClue(prediction int, actualNumber int) (models.Clue, error) {
	var clue models.Clue

	clue.Positive = 0
	clue.Negative = 0

	if prediction == actualNumber {
		clue.Positive = 4
		clue.Negative = 0
		return clue, nil
	}
	preds := strconv.Itoa(prediction)
	actns := strconv.Itoa(actualNumber)
	for i, _ := range preds {
		if preds[i] == actns[i] {
			clue.Positive++
			continue
		}
		for j, _ := range actns {
			if preds[i] == actns[j] {
				clue.Negative++
				continue
			}
		}
	}
	return clue, nil
}
func CalculateWinner() int {
	return 0
}
