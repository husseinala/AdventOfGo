package main

import (
	"fmt"
	"unicode"
)

type numberPosition struct {
	number          string
	y, startX, endX int
}

func dayThree() {
	fileScanner := readFile("day3.txt")

	symbolMap := map[string]bool{}
	gears := map[string][]int{}
	var numbers []numberPosition

	y := 0

	for fileScanner.Scan() {
		num := ""
		x := 0

		for _, ch := range fileScanner.Text() {
			if unicode.IsDigit(ch) {
				num += string(ch)
			} else {
				if num != "" {
					numbers = append(numbers, numberPosition{num, y, x - len(num), x - 1})
					num = ""
				}

				if ch != '.' {
					key := intToStr(y) + ":" + intToStr(x)
					symbolMap[key] = true

					if ch == '*' {
						gears[key] = []int{}
					}
				}
			}
			x++
		}

		if num != "" {
			numbers = append(numbers, numberPosition{num, y, x - len(num), x - 1})
			num = ""
		}

		y++
	}

	dayThreePartOneSum(numbers, symbolMap)
	dayThreePartTwoSum(numbers, gears)

}

func dayThreePartOneSum(numbers []numberPosition, symbolMap map[string]bool) {
	sum := 0

mainLoop:
	for _, n := range numbers {
		for x := n.startX - 1; x <= n.endX+1; x++ {
			if mapContainsKey(symbolMap, intToStr(n.y)+":"+intToStr(x)) ||
				mapContainsKey(symbolMap, intToStr(n.y-1)+":"+intToStr(x)) ||
				mapContainsKey(symbolMap, intToStr(n.y+1)+":"+intToStr(x)) {
				sum += strToInt(n.number)

				continue mainLoop
			}
		}
	}

	fmt.Println("Part One: " + intToStr(sum))
}

func dayThreePartTwoSum(numbers []numberPosition, gears map[string][]int) {
	sum := 0

	for _, n := range numbers {
		for x := n.startX - 1; x <= n.endX+1; x++ {
			for y := n.y - 1; y <= n.y+1; y++ {
				key := intToStr(y) + ":" + intToStr(x)
				if mapContainsKey(gears, key) {
					gears[key] = append(gears[key], strToInt(n.number))
				}
			}
		}
	}

	for _, gear := range gears {
		if len(gear) == 2 {
			sum += gear[0] * gear[1]
		}
	}

	fmt.Println("Part Two: " + intToStr(sum))
}
