package models

type Session struct {
	Id     int    `json:"id"`
	Date   string `json:"date"`
	Turn          int           `json:"turn"`          // Turn Count
	Player1       User          `json:"player1"`       // Player1 Name
	Player2       User          `json:"player2"`       // Player2 Name
	Player1Number int           `json:"player1number"` // Player1 Number
	Player2Number int           `json:"player2number"` // Player2 Number
	Predictor     int           `json:"predictor"`     // Who is Predicting	
	Start  int    `json:"start"`  // Is Game Started or not.0 is not started.1 is player1 ready.2 is player2 ready.3 is both ready.
	End    int    `json:"end"`    // Is Game ended or not.0 or 1.
	Winner int    `json:"winner"` // It indicates the winner.It can be 0(Not Ended),1,2
} 

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Clue struct {
	Positive int `json:"positive"`
	Negative int `json:"negative"`
}

type Move struct {
	Id            int `json:"id"`
	Session       `json:"session"`
	Clue          `json:"clue"` // Clues like +1/-1
	Prediction    int           `json:"prediction"`    // Prediction
	Action        string        `json:"action"`        // Action
}