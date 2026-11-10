package main

import "fmt"

func main() {
	var W, H, R int
	_, err := fmt.Scan(&W, &H, &R)
	if err != nil {
		return
	}

	if W >= 2*R && H >= 2*R {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}
