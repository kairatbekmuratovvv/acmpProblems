package main

import "fmt"

func main() {
	var queue = map[string]string{
		"q": "w", "w": "e", "e": "r", "r": "t", "t": "y",
		"y": "u", "u": "i", "i": "o", "o": "p", "p": "a",
		"a": "s", "s": "d", "d": "f", "f": "g", "g": "h",
		"h": "j", "j": "k", "k": "l", "l": "z", "z": "x",
		"x": "c", "c": "v", "v": "b", "b": "n", "n": "m",
		"m": "q",
	}
	var userInput string
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return
	}

	fmt.Println(queue[userInput])
}
