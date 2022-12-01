package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	elvesSnacksCalories := getElvesSnacksCaloriesList(file)

	fmt.Println(getSumTopThreeCalories(elvesSnacksCalories))

}

/*
We will read elf snacks calories and add it to a list. Sort that list and sum the top 3
*/

func getElvesSnacksCaloriesList(caloriesList *os.File) []int {
	scanner := bufio.NewScanner(caloriesList)

	var totalCalories int

	elfsSnacksCalories := make([]int, 0, 100)

	for scanner.Scan() {
		caloriesString := scanner.Text()
		if caloriesString == "" {
			elfsSnacksCalories = append(elfsSnacksCalories, totalCalories)
			totalCalories = 0
			continue
		}

		actualCalories, err := strconv.Atoi(caloriesString)
		if err != nil {
			panic(err)
		}

		totalCalories += actualCalories
	}

	return elfsSnacksCalories

}

func getSumTopThreeCalories(elfsSnacksCalories []int) int {
	sort.Ints(elfsSnacksCalories)
	topThreeCalories := 0
	listLenght := len(elfsSnacksCalories)

	for i := 1; i < 4; i++ {
		topThreeCalories += elfsSnacksCalories[listLenght-i]
	}

	return topThreeCalories
}
