package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Println(getTotalAssignmentPairs(file))
}

func getTotalAssignmentPairs(inputFile *os.File) int {
	scanner := bufio.NewScanner(inputFile)

	totalAssignmentPairs := 0

	for scanner.Scan() {
		sectionAssignmentPair := strings.Split(scanner.Text(), ",")
		firstPair := sectionAssignmentPair[0]
		secondPair := sectionAssignmentPair[1]

		firstAssignedSections := strings.Split(firstPair, "-")
		secondAssignedSections := strings.Split(secondPair, "-")

		firstAssignedFirstNumber, _ := strconv.Atoi(firstAssignedSections[0])
		firstAssignedSecondNumber, _ := strconv.Atoi(firstAssignedSections[1])
		secondAssignedFirstNumber, _ := strconv.Atoi(secondAssignedSections[0])
		secondAssignedSecondNumber, _ := strconv.Atoi(secondAssignedSections[1])

		if firstAssignedFirstNumber <= secondAssignedSecondNumber && firstAssignedSecondNumber >= secondAssignedFirstNumber {
			totalAssignmentPairs += 1
			continue
		}

		if secondAssignedFirstNumber <= firstAssignedFirstNumber && secondAssignedSecondNumber >= firstAssignedFirstNumber {
			totalAssignmentPairs += 1
		}
	}

	return totalAssignmentPairs
}
