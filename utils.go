package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(name string) *bufio.Scanner {
	readFile, err := os.Open("resources/" + name)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func mapContainesKey[K comparable, V any](testMap map[K]V, key K) bool {
	_, ok := testMap[key]

	return ok
}

func strToInt(str string) int {
	val, _ := strconv.Atoi(str)

	return val
}
