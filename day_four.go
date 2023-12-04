package main

import (
	"fmt"
	"strings"
)

func dayFour() {
	dayFourPartOne()
	dayFourPartTwo()
}

func dayFourPartOne() {
	sum := 0

	f := func(id int, winningNumbers map[string]bool, cardNumbers []string) {
		cardScore := 0

		for _, n := range cardNumbers {
			if mapContainsKey(winningNumbers, n) {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore *= 2
				}
			}
		}

		sum += cardScore
	}

	processDayFourInput(f)

	fmt.Println("Part 1: " + intToStr(sum))
}

func dayFourPartTwo() {
	sum := 0
	numberOfCards := map[int]int{}

	f := func(id int, winningNumbers map[string]bool, cardNumbers []string) {
		if !mapContainsKey(numberOfCards, id) {
			numberOfCards[id] = 1
		}

		winIncrement := numberOfCards[id]
		winningNumber := 0

		for _, n := range cardNumbers {
			if mapContainsKey(winningNumbers, n) {
				winningNumber += 1

				if mapContainsKey(numberOfCards, id+winningNumber) {
					numberOfCards[id+winningNumber] = numberOfCards[id+winningNumber] + winIncrement
				} else {
					numberOfCards[id+winningNumber] = 1 + winIncrement
				}
			}
		}

		sum += winIncrement
	}

	processDayFourInput(f)

	fmt.Println("Part 2: " + intToStr(sum))
}

func processDayFourInput(action func(id int, winningNumbers map[string]bool, cardNumbers []string)) {
	fileScanner := readFile("day4.txt")

	for fileScanner.Scan() {
		id, numbers := destructure(strings.Split(fileScanner.Text(), ":"))
		cardId := strToInt(strings.TrimSpace(id[4:]))

		winningNumbers, cardNumbers := destructure(strings.Split(numbers, "|"))

		winningNumbersDic := map[string]bool{}

		for _, n := range strings.Split(strings.TrimSpace(winningNumbers), " ") {
			if strings.TrimSpace(n) != "" {
				winningNumbersDic[n] = true
			}
		}

		action(cardId, winningNumbersDic, strings.Split(strings.TrimSpace(cardNumbers), " "))
	}
}
