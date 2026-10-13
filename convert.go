package main

import (
    "fmt"
    "strings"
)

func main() {
    const Usd = 83.9
    const Euro = 94.0

    var amount float64
    var choice int
    var answer string
    var trash string 

    for {
        fmt.Println("\n1: USD в RUB")
        fmt.Println("2: EUR в RUB")
        fmt.Println("3: RUB в USD")
        fmt.Println("4: RUB в EUR")
        fmt.Print("Выберите вариант (1-4): ")

        _, err := fmt.Scan(&choice)
        if err != nil || choice < 1 || choice > 4 {
            fmt.Println("Ошибка: нужно ввести число от 1 до 4!")
            fmt.Scan(&trash) 
            continue
        }

        fmt.Print("Введите сумму для конвертации: ")
        _, err = fmt.Scan(&amount)
        if err != nil {
            fmt.Println("Ошибка: нужно ввести число!")
            fmt.Scan(&trash) 
            continue
        }

        if choice == 1 {
            fmt.Printf("Результат: %.2f RUB\n", amount*Usd)
        } else if choice == 2 {
            fmt.Printf("Результат: %.2f RUB\n", amount*Euro)
        } else if choice == 3 {
            fmt.Printf("Результат: %.2f USD\n", amount/Usd)
        } else if choice == 4 {
            fmt.Printf("Результат: %.2f EUR\n", amount/Euro)
        }

        fmt.Print("Сконвертировать еще yes/no ? ")
        fmt.Scan(&answer)
        answerLower := strings.ToLower(answer)
        if answerLower == "n" || answerLower == "no" {
            fmt.Println("Программа завершена")
            break
        }
    }
}