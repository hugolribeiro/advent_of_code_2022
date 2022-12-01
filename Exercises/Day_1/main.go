package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(getHighestCalories(file))

	defer file.Close()

}

func getHighestCalories(caloriesList *os.File) int {
	scanner := bufio.NewScanner(caloriesList)

	totalCalories := 0

	totalCalories, highestCalories := 0, 0

	for scanner.Scan() {
		caloriesString := scanner.Text()
		if caloriesString == "" {
			if totalCalories > highestCalories {
				highestCalories = totalCalories
			}

			totalCalories = 0
			continue
		}

		actualCalories, err := strconv.Atoi(caloriesString)
		if err != nil {
			panic(err)
		}

		totalCalories += actualCalories
	}

	return highestCalories
}
