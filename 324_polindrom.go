package main

import "fmt"

func main() {
	var number int
	_, err := fmt.Scan(&number)

	if err != nil {
		return
	}
	var elements = [4]int{}
	elements[0] = number % 10
	elements[1] = number % 100 / 10
	elements[2] = number % 1000 / 100
	elements[3] = number % 10000 / 1000

	if elements[0] == elements[3] && elements[1] == elements[2] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
