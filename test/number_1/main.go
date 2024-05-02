package main

import (
	"fmt"
)

func differentTipesOfComputers(num int) string {
	if num >= 11 && num <= 19 {
		return fmt.Sprintf("%d компьютеров", num)
	} else {
		lastNumber := num % 10
		switch lastNumber {
		case 1:
			return fmt.Sprintf("%d компьютер", num)
		case 2, 3, 4:
			return fmt.Sprintf("%d компьютера", num)
		default:
			return fmt.Sprintf("%d компьютеров", num)
		}
	}
}

func main() {
	fmt.Println(differentTipesOfComputers(1201))
	fmt.Println(differentTipesOfComputers(41))
	fmt.Println(differentTipesOfComputers(1048))
}
