package main

import "fmt"

func main() {
	var guests int

	_, err := fmt.Scan(&guests)

	if err != nil {
		return
	}

	if guests == 1 {
		fmt.Println(0)
	} else if guests%2 == 0 {
		fmt.Println(guests / 2)
	} else {
		fmt.Println(guests)
	}

}
