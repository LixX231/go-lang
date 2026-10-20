package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const Usd = 83.9
	const Euro = 94.0

	http.HandleFunc("/convertor", func(a http.ResponseWriter, r *http.Request) {
		a.Header().Set("Content-Type", "text/html; charset=utf-8")

		amountStr := r.FormValue("amount")
		if amountStr == "" {
			amountStr = "100"
		}
		var amount float64
		fmt.Sscanf(amountStr, "%f", &amount)

		result := "Выберите действие"

		switch r.FormValue("action") {
		case "usd_rub":
			result = fmt.Sprintf("%.2f USD = %.2f RUB", amount, amount*Usd)
		case "eur_rub":
			result = fmt.Sprintf("%.2f EUR = %.2f RUB", amount, amount*Euro)
		case "rub_usd":
			result = fmt.Sprintf("%.2f RUB = %.2f USD", amount, amount/Usd)
		case "rub_eur":
			result = fmt.Sprintf("%.2f RUB = %.2f EUR", amount, amount/Euro)
		}

		fmt.Fprintf(a, `
			<h1>Конвертер Валют</h1>
			<form method="GET" action="/convertor">
				<p>Сумма: <input type="number" name="amount" value="%s" step="any" style="padding: 5px;"></p>
				<p>
					<button type="submit" name="action" value="usd_rub">USD в RUB</button>
					<button type="submit" name="action" value="eur_rub">EUR в RUB</button>
					<button type="submit" name="action" value="rub_usd">RUB в USD</button>
					<button type="submit" name="action" value="rub_eur">RUB в EUR</button>
				</p>
			</form>
			<h2>Результат: %s</h2>
			<br><br>
			<p><a href="/generator">Перейти в Дженератор 3000</a> | <a href="/gleb">Перейти к Айфону Глеба</a></p>
		`, amountStr, result)
	})

	http.HandleFunc("/generator", func(a http.ResponseWriter, r *http.Request) {
		a.Header().Set("Content-Type", "text/html; charset=utf-8")

		randomNumber := ((time.Now().UnixNano() / 1000) % 100) + 1

		fmt.Fprintf(a, `
	<h1>Дженератор чисел 3000</h1>
	<p>Тут чилл</p>
	<p>Ваше случайное число: <strong>%d</strong></p>
	
	<!-- Кнопка, которая просто обновляет текущую страницу -->
	<form action="/generator" method="GET">
		<button type="submit" style="padding: 10px 20px; font-size: 16px; cursor: pointer;">
			Сгенерировать новое число
		</button>
	</form>
	
	<br><br>
	<a href="/convertor">На главную (к конвертеру)</a>
	`, randomNumber)
	})

	http.HandleFunc("/gleb", func(a http.ResponseWriter, r *http.Request) {
		a.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(a, `
			<h1>Айфон Глеба</h1>
			<p>Здесь все тоже чилл.</p>
			<br>
			<a href="/convertor">На главную (к конвертеру)</a>
		`)
	})

	fmt.Println("Запущен на http://localhost:8080/convertor")
	http.ListenAndServe(":8080", nil)
}
