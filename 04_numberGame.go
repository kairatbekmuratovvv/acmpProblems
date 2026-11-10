package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findRemainder(userNumber int64, localNumber int64) int64 {
	remainder := localNumber - userNumber
	return remainder
}

func main() {
	var userNumber, localNumber, remainder int64
	localNumber = 9

	_, err := fmt.Scan(&userNumber)
	if err != nil {
		return
	}
	remainder = findRemainder(userNumber, localNumber)
	total := [3]int64{userNumber, localNumber, remainder}
	var strSlice []string
	for _, num := range total {
		strSlice = append(strSlice, strconv.FormatInt(num, 10))
	}
	combinedStr := strings.Join(strSlice, "")
	finalNumber, _ := strconv.ParseInt(combinedStr, 10, 64)
	fmt.Println(finalNumber)
}
