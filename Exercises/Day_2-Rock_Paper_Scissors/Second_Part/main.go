package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var roundScores = map[string]int{
	"Won":  6,
	"Draw": 3,
	"Lost": 0,
}

var shapesScores = map[string]int{
	"Rock":    1,
	"Paper":   2,
	"Scissor": 3,
}

const (
	draw    = "Draw"
	lost    = "Lost"
	won     = "Won"
	rock    = "Rock"
	paper   = "Paper"
	scissor = "Scissor"
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

	totalScore := 0

	for scanner.Scan() {
		roundInput := strings.Split(scanner.Text(), " ")
		opponentChoice, roundResult := roundInput[0], roundInput[1]

		totalScore += getRoundScore(opponentChoice, roundResult)
	}

	return totalScore
}

func getRoundScore(opponentChoice, roundResult string) (roundScore int) {
	if opponentChoice == "A" {
		roundScore = getScoreToRockChoice(roundResult)
	} else if opponentChoice == "B" {
		roundScore = getScoreToPaperChoice(roundResult)
	} else if opponentChoice == "C" {
		roundScore = getScoreToScissorChoice(roundResult)
	}

	return
}

func getScoreToRockChoice(roundResult string) (roundScore int) {
	if roundResult == "X" {
		roundScore += roundScores[lost]
		roundScore += shapesScores[scissor]
	} else if roundResult == "Y" {
		roundScore += roundScores[draw]
		roundScore += shapesScores[rock]
	} else if roundResult == "Z" {
		roundScore += roundScores[won]
		roundScore += shapesScores[paper]
	}

	return
}

func getScoreToPaperChoice(roundResult string) (roundScore int) {
	if roundResult == "X" {
		roundScore += roundScores[lost]
		roundScore += shapesScores[rock]
	} else if roundResult == "Y" {
		roundScore += roundScores[draw]
		roundScore += shapesScores[paper]
	} else if roundResult == "Z" {
		roundScore += roundScores[won]
		roundScore += shapesScores[scissor]
	}

	return
}

func getScoreToScissorChoice(roundResult string) (roundScore int) {
	if roundResult == "X" {
		roundScore += roundScores[lost]
		roundScore += shapesScores[paper]
	} else if roundResult == "Y" {
		roundScore += roundScores[draw]
		roundScore += shapesScores[scissor]
	} else if roundResult == "Z" {
		roundScore += roundScores[won]
		roundScore += shapesScores[rock]
	}

	return
}
