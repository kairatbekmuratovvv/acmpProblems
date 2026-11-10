package main

import "fmt"

func main() {
	var a1, a2, a3 int
	_, err := fmt.Scan(&a1, &a2, &a3)
	if err != nil {
		return
	}

	if a1+a2 == a3 || a1+a3 == a2 || a2+a3 == a1 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
