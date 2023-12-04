package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	dayFour()

	fmt.Println("Run Time: " + strconv.FormatInt(time.Now().UnixMilli()-startTime, 10))
}
