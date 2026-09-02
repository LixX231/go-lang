package main

import "fmt"

func main() {
	var number int = 122

	if number > 0 {
		var first int = number / 100
		var second int = (number / 10) % 10
		var third int = number % 10
		if first != second && first != third && second != third {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		fmt.Println("NO", number)
	}

}