package main

import (
	"fmt"
	"math"
)

func main() {

	var user_number int

	_, err := fmt.Scan(&user_number)

	if err != nil {
		return
	}

	a := 1
	for i := 2; i < user_number; i++ {
		if user_number%i == 0 {
			a += 1
		}
	}
	base := 2
	pow := int(math.Pow(float64(base), float64(a)))

	if user_number == 2 && user_number > 0 {
		fmt.Println("YES")
	} else if user_number == pow && user_number > 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
