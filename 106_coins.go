package main

import "fmt"

func main() {
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}
	var t_count int
	var g_count int

	for i := 0; i < count; i++ {
		var a int
		_, err := fmt.Scan(&a)
		if err != nil {
			return
		}
		if a == 0 {
			t_count++
		} else {
			g_count++
		}
	}
	if t_count > g_count {
		fmt.Print(g_count)
	} else {
		fmt.Print(t_count)
	}
}
