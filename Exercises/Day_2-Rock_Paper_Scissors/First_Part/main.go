package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var shapesScores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var roundScores = map[string]int{
	"Won":  6,
	"Draw": 3,
	"Lost": 0,
}

const (
	draw = "Draw"
	lost = "Lost"
	won  = "Won"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	totalScore := getTotalScoreFromFile(file)

	fmt.Println(totalScore)
}

func getTotalScoreFromFile(guideStrategy *os.File) int {
	scanner := bufio.NewScanner(guideStrategy)

	var opponentChoice, yourChoice string

	totalScore := 0

	for scanner.Scan() {
		bothChoices := strings.Split(scanner.Text(), " ")
		opponentChoice, yourChoice = bothChoices[0], bothChoices[1]

		totalScore += getRoundScore(opponentChoice, yourChoice)
	}

	return totalScore
}

func getRoundScore(opponentChoice, yourChoice string) (yourTotalScore int) {
	yourTotalScore = shapesScores[yourChoice]

	// Opponent choice Rock statement
	if opponentChoice == "A" {
		if yourChoice == "X" {
			yourTotalScore += roundScores[draw]
			return
		}
		if yourChoice == "Y" {
			yourTotalScore += roundScores[won]
			return
		}
	}

	// Opponent choice Paper statement
	if opponentChoice == "B" {
		if yourChoice == "Y" {
			yourTotalScore += roundScores[draw]
			return
		}
		if yourChoice == "Z" {
			yourTotalScore += roundScores[won]
			return
		}
	}

	// Opponent choice Scissor statement
	if opponentChoice == "C" {
		if yourChoice == "X" {
			yourTotalScore += roundScores[won]
			return
		}
		if yourChoice == "Z" {
			yourTotalScore += roundScores[draw]
			return
		}
	}

	yourTotalScore += roundScores[lost]
	return
}
