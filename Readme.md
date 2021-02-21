# Bulls and Cows Game
This project is a number game that plays with 2 people. You can create a game session and after that invite your friend with session id.

## Rules
* Each player must define a number with specified digit.The first number can not be 0.
* Every turn, each player must make a prediction about opponent's number by order.
* After every prediction, game gives a clue about opponent's number. This clue consist only a positive integer and/or negative integer.
* If numbers are match but digit locations are different it gives negative and if numbers match and digit locations are true it gives positive. If neither digits nor digit locations are true it gives 0.
* After every turn players must calculate opponent's number and make better prediction.
* The game is continues until one of the players guess  opponent's number right.
* The winner player win points and game goes to a new round.<br>
Further Information: https://en.wikipedia.org/wiki/Bulls_and_Cows

## Example Game
* Number is 1234.
* Prediction: 5678
* Game System Clue: 0
* Prediction: 4321
* Game System Clue: -4
* Prediction: 1243
* Game System Clue: +2-2
* Prediction:1234
* Game System Clue: +4 - WIN

## Scoring System
TODO: will be implemented

