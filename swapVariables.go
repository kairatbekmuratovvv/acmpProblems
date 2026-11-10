package main

import "fmt"

func main() {

	a := 5
	b := 6

	a, b = b, a

	fmt.Println(a, b)

}
