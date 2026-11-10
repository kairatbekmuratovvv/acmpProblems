package main

import "fmt"

func main() {

	var total int

	_, err := fmt.Scan(&total)

	if err != nil {
		return
	}

	a := total / 6

	fmt.Println(a, a*4, a)

}
