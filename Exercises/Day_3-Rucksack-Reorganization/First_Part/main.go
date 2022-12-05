package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Println(getSumPriorities(file))
}

func getSumPriorities(file *os.File) int {
	scanner := bufio.NewScanner(file)

	totalPriorityItems := 0

	for scanner.Scan() {
		rucksackContent := scanner.Text()
		firstCompartment, secondCompartment := rucksackContent[0:len(rucksackContent)/2], rucksackContent[len(rucksackContent)/2:]
		letterInCommon := getLetterInCommon(firstCompartment, secondCompartment)
		itemPriority := getValueOfLetter(letterInCommon)

		totalPriorityItems += itemPriority
	}
	return totalPriorityItems
}

func getLetterInCommon(firstString, secondString string) string {
	letterInCommon := ""

	for _, char := range firstString {
		if strings.Contains(secondString, string(char)) {
			letterInCommon = string(char)
			break
		}
	}
	return letterInCommon
}

func getValueOfLetter(letter string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(alphabet, letter) + 1
}
