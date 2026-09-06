package main

import "fmt"

func main() {
	var ticket string = "385934"

	var sum1 = ticket[0] + ticket[1] + ticket[2]
	var sum2 = ticket[3] + ticket[4] + ticket[5]

	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
