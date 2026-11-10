package main

import "fmt"

func main() {
	var s1, s2, s3 int

	_, err := fmt.Scan(&s1, &s2, &s3)

	if err != nil {
		return
	}

	maxSalary := max(s1, s2, s3)
	minSalary := min(s1, s2, s3)

	fmt.Println(maxSalary - minSalary)

}
