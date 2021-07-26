package logic

import (
	"errors"
	"strconv"

	models "github.com/anilkusc/BullsAndCows/models"
)

func IsNumberLegal(number int) error {
	if number < 1000 || number > 9999 {
		return errors.New("Number does not have 4 digits.")
	}
	numbers := strconv.Itoa(number)
	// TODO: optimize this control
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			if numbers[i] == numbers[j] {
				return errors.New("Duplicated Digits")
			}
		}
	}
	return nil

}
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
