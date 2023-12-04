package main

import (
	"bufio"
	"fmt"
	"strings"
)

func dayTwo() {
	dayTwoPartTwo(readFile("day2.txt"))
}

func dayTwoPartOne(fileScanner *bufio.Scanner) {
	maxCubes := map[string]int{"red": 12, "green": 13, "blue": 14}
	sum := 0

	for fileScanner.Scan() {
		id, gems := destructure(strings.Split(fileScanner.Text(), ":"))
		gameNumber := strToInt(id[5:])

		sets := strings.Split(gems, ";")

		isPossible := true

	mainLoop:
		for _, set := range sets {
			setGems := strings.Split(set, ",")

			for _, gem := range setGems {
				valueKey := strings.Split(strings.TrimSpace(gem), " ")

				if maxCubes[valueKey[1]] < strToInt(valueKey[0]) {
					isPossible = false
					break mainLoop
				}
			}
		}

		if isPossible {
			sum += gameNumber
		}
	}

	fmt.Println(sum)
}

func dayTwoPartTwo(fileScanner *bufio.Scanner) {
	sum := 0

	for fileScanner.Scan() {
		_, gems := destructure(strings.Split(fileScanner.Text(), ":"))

		sets := strings.Split(gems, ";")

		minCubes := map[string]int{}
		for _, set := range sets {
			setGems := strings.Split(set, ",")

			for _, gem := range setGems {
				valueKey := strings.Split(strings.TrimSpace(gem), " ")
				key := valueKey[1]
				value := strToInt(valueKey[0])

				if !(mapContainsKey(minCubes, key) && minCubes[key] > value) {
					minCubes[key] = value
				}
			}
		}

		setsSum := 1

		for _, value := range minCubes {
			setsSum = setsSum * value
		}

		sum += setsSum
	}

	fmt.Println(sum)
}
