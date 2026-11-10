package main

import "fmt"

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b, &c)

	if err != nil {
		return
	}

	var S = 2 * (a*b + b*c + a*c)
	var V = a * b * c

	fmt.Println(S, V)

}
