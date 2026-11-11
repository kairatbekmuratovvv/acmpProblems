package main

import "fmt"

func main() {
	var L, W, H int
	_, err := fmt.Scan(&L, &W, &H)
	if err != nil {
		return
	}

	var S = 2 * (L + W) * H
	var N = (S + 15) / 16
	fmt.Println(N)
}
