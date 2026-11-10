package main

import "fmt"

func main() {

	var A1, A2, A3, B1, B2, B3 int

	_, err := fmt.Scan(&A1, &A2, &A3, &B1, &B2, &B3)

	if err != nil {
		return
	}

	var value = [3]int{B1, B2, B3}
	var gold = [3]int{A1, A2, A3}

	for i := len(value); i > 0; i-- {
		for j := 1; j < i; j++ {
			if value[j-1] > value[j] {
				intermediate := value[j]
				value[j] = value[j-1]
				value[j-1] = intermediate
			}

		}
	}

	for i := len(gold); i > 0; i-- {
		for j := 1; j < i; j++ {
			if gold[j-1] > gold[j] {
				gold[j-1], gold[j] = gold[j], gold[j-1]
			}
		}
	}

	var total = map[int]int{}

	for i := 0; i < 3; i++ {
		total[i] = value[i] * gold[i]
	}
	var result int

	for _, value := range total {
		result += value
	}
	fmt.Println(result)
}
