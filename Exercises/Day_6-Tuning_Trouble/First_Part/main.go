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
	for i := 0; i < len(deviceSubroutineMessage)-4; i++ {
		fourCharacters := deviceSubroutineMessage[i : i+4]
		if hasUniqueLetters(fourCharacters) {
			return i + 4
		}
	}
	return 0
}

func hasUniqueLetters(fourCharacters string) bool {
	for i := 0; i <= 2; i++ {
		for j := i + 1; j <= 3; j++ {
			if fourCharacters[i] == fourCharacters[j] {
				return false
			}
		}
	}
	return true
}
