package utils

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json: "message"`
	ComputerChoice    string `json: "computer_choice"`
	RoundResult       string `json: "round_result"`
	ComputerChoiceInt int    `json: "computer_choice_int"`
	ComputerScore     string `json: "computer_score"`
	PlayerScore       string `json: "player_score"`
}

var winMessages = []string{
	"Good job !",
	"Well done !",
}

var loseMessages = []string{
	"Oh wow...",
	"Try again !",
}

var drawMessages = []string{
	"you both tie!",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "The computer chose rock"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "The computer chose paper"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "The computer chose scissors"
	}

	messageInt := rand.Intn(2)
	var message string

	if playerValue == computerValue {
		roundResult = "it's a tie"
		message = drawMessages[0]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "Player win !"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "Computer win !"
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
