package main

import "fmt"

func main() {
	var number int = 0

	if number > 0 {
		fmt.Println("Число положительное", number)
	} else if number < 0 {
		fmt.Println("Число отрицательное", number)
	} else {
		fmt.Println("Ноль", number)
	}
}