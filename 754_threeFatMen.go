package main

import "fmt"

func main() {

	var m1, m2, m3 int

	_, err := fmt.Scan(&m1, &m2, &m3)

	if err != nil {
		return
	}

	var max1 int
	var min1 int

	if m1 >= m2 {
		max1 = m1
		min1 = m2
	} else {
		max1 = m2
		min1 = m1
	}

	if max1 > m3 && min1 < m3 {

	} else if max1 >= m3 && min1 >= m3 {
		min1 = m3
	} else {
		max1 = m3
	}

	if max1 <= 727 && min1 >= 94 {
		fmt.Print(max1)
	} else {
		fmt.Print("Error")
	}

}
