package main

import "fmt"

func formula(m int, n int) int {
	return (m - 1) * (n - 1)
}

func main() {
	var m, n int
	_, err := fmt.Scan(&m, &n)

	if err != nil {
		return
	}
	result := formula(m, n)
	fmt.Println(result)

}
