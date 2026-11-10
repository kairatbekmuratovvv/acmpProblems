package main

import "fmt"

func main() {
	var month = map[int]string{
		1:  "Winter",
		2:  "Winter",
		3:  "Spring",
		4:  "Spring",
		5:  "Spring",
		6:  "Summer",
		7:  "Summer",
		8:  "Summer",
		9:  "Autumn",
		10: "Autumn",
		11: "Autumn",
		12: "Winter",
	}
	var userInput int
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return
	}

	if val, ok := month[userInput]; ok {
		fmt.Print(val)
	} else {
		fmt.Print("Error")
	}
}
