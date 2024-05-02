package main

import (
	"fmt"
)

func isSimple(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5

	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}

func findSimple(minNum, maxNum int) []int {
	var simple []int
	for num := minNum; num <= maxNum; num++ {
		if isSimple(num) {
			simple = append(simple, num)
		}
	}

	return simple
}

func main() {
	minNum := 11
	maxNum := 20
	result := findSimple(minNum, maxNum)
	fmt.Println(result)
}
