package main

import (
	"fmt"
)

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}
	var values []int

	for i := 0; i < count; i++ {
		var a int
		_, err := fmt.Scan(&a)
		if err != nil {
			return
		}
		values = append(values, a)

	}
	reverseSlice(values)

	args := make([]any, len(values))
	for i, v := range values {
		args[i] = v
	}
	fmt.Println(args...)

}
