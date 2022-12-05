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

func getSumPriorities(inputFile *os.File) int {
	scanner := bufio.NewScanner(inputFile)
	totalPriorities := 0

	for scanner.Scan() {
		firstElfOfGroup := scanner.Text()
		scanner.Scan()
		secondElfOfGroup := scanner.Text()
		scanner.Scan()
		thirdElfOfGroup := scanner.Text()

		commonItem := getCommonLetter(firstElfOfGroup, secondElfOfGroup, thirdElfOfGroup)
		totalPriorities += getItemPriority(commonItem)
	}

	return totalPriorities
}

func getCommonLetter(firstString, secondString, thirdString string) string {
	letterInCommon := ""
	for _, char := range firstString {
		if strings.Contains(secondString, string(char)) && strings.Contains(thirdString, string(char)) {
			letterInCommon = string(char)
			break
		}
	}
	return letterInCommon
}

func getItemPriority(letter string) int {
	possibleItems := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(possibleItems, letter) + 1
}
