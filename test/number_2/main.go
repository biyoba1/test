package main

import (
	"fmt"
)

func findAllDelitels(numbers []int) []int {
	minimal := numbers[0]
	for _, num := range numbers {
		if num < minimal {
			minimal = num
		}
	}

	delitel := make([]int, 0)

	for i := 1; i <= minimal; i++ {
		isDelitel := true
		for _, num := range numbers {
			if num%i != 0 {
				isDelitel = false
				break
			}
		}

		if isDelitel {
			delitel = append(delitel, i)
		}
	}

	return delitel
}

func main() {
	numbers := []int{42, 12, 18}
	result := findAllDelitels(numbers)
	fmt.Println(result)
}
