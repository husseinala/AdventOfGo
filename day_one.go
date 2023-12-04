package main

import (
	"fmt"
	"unicode"
)

func dayOne() {
	digitNumbers := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	fileScanner := readFile("day1.txt")

	sum := 0

	for fileScanner.Scan() {
		sum += getDigits(fileScanner.Text(), digitNumbers)
	}

	fmt.Println("" + intToStr(sum))
}

func getDigits(line string, digitNumbers map[string]int) int {
	var firstDigit int
	var lastDigit int

	for index, ch := range line {
		var num int

		if unicode.IsDigit(ch) {
			num = int(ch - '0')
		} else if (len(line)-index) >= 3 && mapContainsKey(digitNumbers, line[index:index+3]) {
			num = digitNumbers[line[index:index+3]]
		} else if (len(line)-index) >= 4 && mapContainsKey(digitNumbers, line[index:index+4]) {
			num = digitNumbers[line[index:index+4]]
		} else if (len(line)-index) >= 5 && mapContainsKey(digitNumbers, line[index:index+5]) {
			num = digitNumbers[line[index:index+5]]
		}

		if num != 0 {
			if firstDigit == 0 {
				firstDigit = num
			}

			lastDigit = num
		}
	}

	return firstDigit*10 + lastDigit
}
