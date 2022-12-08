package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	deviceSubroutineMessage := scanner.Text()

	firstMarker := getFirstMarkerIndex(deviceSubroutineMessage)

	fmt.Println(firstMarker)
}

func getFirstMarkerIndex(deviceSubroutineMessage string) int {
	for i := 0; i < len(deviceSubroutineMessage)-14; i++ {
		fourteenCharacters := deviceSubroutineMessage[i : i+14]
		if hasUniqueLetters(fourteenCharacters) {
			return i + 14
		}
	}
	return 0
}

func hasUniqueLetters(fourteenCharacters string) bool {
	for i := 0; i <= 12; i++ {
		for j := i + 1; j <= 13; j++ {
			if fourteenCharacters[i] == fourteenCharacters[j] {
				return false
			}
		}
	}
	return true
}
