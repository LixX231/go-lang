
package main

import "fmt"

func main() {
	var imt float64
	var weight float64 = 80
	var height float64 = 1.88

	imt = weight / (height * height)

	fmt.Println("ИМТ:", imt)
}