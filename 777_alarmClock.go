package main

import "fmt"

func main() {
	const M = 12
	var S, T int

	_, err := fmt.Scan(&S, &T)
	if err != nil {
		return
	}

	var d = T - S
	if d < 0 {
		d = (M - S) + T
	}
	fmt.Println(d)

}
