package main

import "fmt"

//Задача 1.1 Объявить переменные всех типов 3 способами
func main() {
	var a int = 10
	var b string = "Ivan"
	var c bool
	var d float64 = 10.5
	fmt.Printf("Формат %T\n", a)
	fmt.Printf("Формат %T\n", b)
	fmt.Printf("Формат %T\n", c)
	fmt.Printf("Формат %T\n", d)
}

func second() {
	var (
		a2 = 10
		b2 = "Ivan"
		c2 = false
		d2 = 10.5
	)
	fmt.Printf("Формат %T", a2)
	fmt.Printf("Формат %T", b2)
	fmt.Printf("Формат %T", c2)
	fmt.Printf("Формат %T", d2)
}

func asdwq() {
	a3 := 10
	b3 := "Ivan"
	c3 := false
	d3 := 10.5
	fmt.Printf("Формат %T", a3)
	fmt.Printf("Формат %T", b3)
	fmt.Printf("Формат %T", c3)
	fmt.Printf("Формат %T", d3)
}

//Задача 1.2 Поменять местами значение 2х переменных
//С использованием 3
func task21() {
	a := 5
	b := 10

	temp := a
	a = b
	b = temp

	fmt.Printf("a = %v, b = %v\n", a, b)
}

//Без использования 3
func task22() {
	a := 5
	b := 10

	a, b = b, a

	fmt.Printf("a = %v, b = %v\n", a, b)
}



/*initioal