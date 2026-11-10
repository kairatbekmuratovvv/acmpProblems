package main

import "fmt"

func calculate(c int, h int, o int) int {
	min := 0
	C := c / 2
	H := h / 6
	O := o / 1

	if C < H {
		min = C
	} else {
		min = H
	}

	if min > O {
		min = O
	}

	return min

}

func main() {
	var C, H, O int

	_, err := fmt.Scan(&C, &H, &O)

	if err != nil {
		return
	}

	result := calculate(C, H, O)
	fmt.Println(result)

}
