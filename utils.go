package main

import (
	"bufio"
	"os"
	"strconv"
)

func readFile(name string) *bufio.Scanner {
	readFile, err := os.Open("resources/" + name)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func mapContainsKey[K comparable, V any](testMap map[K]V, key K) bool {
	_, ok := testMap[key]

	return ok
}

func strToInt(str string) int {
	val, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return val
}

func intToStr(val int) string {
	return strconv.Itoa(val)
}

func destructure[V any](list []V) (V, V) {
	return list[0], list[1]
}

func destructure3[V any](list []V) (V, V, V) {
	return list[0], list[1], list[2]
}
