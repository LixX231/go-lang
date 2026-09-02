package main

import "fmt"

func main() {
	var number int = 322

	if number > 0 {
		var first int = number / 100
		fmt.Println("Цифра 1:", first)
	}
}
