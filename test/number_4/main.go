package main

import "fmt"

func table(n int) {
	fmt.Print("    ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Printf("%d |", i)
		for j := 1; j <= n; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}

func main() {
	n := 6
	table(n)
}
