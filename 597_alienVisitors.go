package main

import "fmt"

func main() {
	var r1, r2, r3 int

	_, err := fmt.Scan(&r1, &r2, &r3)

	if err != nil {
		return
	}
	total := r2 + r3

	if r1 >= total {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}
