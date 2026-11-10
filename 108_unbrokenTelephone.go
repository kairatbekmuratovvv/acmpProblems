package main

import "fmt"

func main() {
	var userInput int
	_, err := fmt.Scan(&userInput)

	if err != nil {
		return
	}

	fmt.Println(userInput)
}
