package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var startingStacks = getStartingStacks()

func getStartingStacks() (stacks [][]string) {
	stacks = append(stacks, []string{
		"S", "M", "R", "N", "W", "J", "V", "T",
	})
	stacks = append(stacks, []string{
		"B", "W", "D", "J", "Q", "P", "C", "V",
	})
	stacks = append(stacks, []string{
		"B", "J", "F", "H", "D", "R", "P",
	})
	stacks = append(stacks, []string{
		"F", "R", "P", "B", "M", "N", "D",
	})
	stacks = append(stacks, []string{
		"H", "V", "R", "P", "T", "B",
	})
	stacks = append(stacks, []string{
		"C", "B", "P", "T",
	})
	stacks = append(stacks, []string{
		"B", "J", "R", "P", "L",
	})
	stacks = append(stacks, []string{
		"N", "C", "S", "L", "T", "Z", "B", "W",
	})
	stacks = append(stacks, []string{
		"L", "S", "G",
	})

	return stacks
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	rearangeCrates(file)

	fmt.Println(getOnlyTopCrates())

}

func rearangeCrates(rearrangementProcedure *os.File) {
	scanner := bufio.NewScanner(rearrangementProcedure)

	for scanner.Scan() {
		text := scanner.Text()

		moves := string(text[5])

		if _, err := strconv.Atoi(string(text[6])); err == nil {
			moves += string(text[6])
		}

		amountMoves, _ := strconv.Atoi(moves)

		indexFrom := strings.Index(text, "m ")
		moveFrom, _ := strconv.Atoi(string(text[indexFrom+2]))

		moveTo, _ := strconv.Atoi(string(text[indexFrom+7]))

		doTheMovement(amountMoves, moveFrom-1, moveTo-1)

	}
}

func doTheMovement(moves, from, to int) {
	for movement := 0; movement < moves; movement++ {
		startingStacks[to] = append(startingStacks[to], startingStacks[from][len(startingStacks[from])-1])
		startingStacks[from] = startingStacks[from][0 : len(startingStacks[from])-1]
	}
}

func getOnlyTopCrates() string {
	message := ""
	for i := 0; i < 9; i++ {
		message += startingStacks[i][len(startingStacks[i])-1]
	}
	return message
}
