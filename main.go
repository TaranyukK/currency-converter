package main

import "fmt"

const (
	USDTOEUR float64 = 0.85
	USDTORUB float64 = 0.013
	EURTORUB         = USDTOEUR * USDTORUB
)

var (
	amount    int
	stockCur  string
	targetCur string
)

func main() {
	amount, stockCur, targetCur := getUserInput()

	resultNum := convertCurrency(amount, stockCur, targetCur)

	fmt.Printf("Количество переведенной валюты из %v в %v: %.2f", stockCur, targetCur, resultNum)
}

func getUserInput() (int, string, string) {
	fmt.Print("Введите валюту, которую хотите конвертировать (USD, EUR, RUB): ")
	getStockCur()
	fmt.Print("Введите количество валюты, которую хотите ковертировать: ")
	getAmount()
	fmt.Print("Введите валюту, в которую хотите конвертировать (USD, EUR, RUB): ")
	getTargetCur()

	return amount, stockCur, targetCur
}

func convertCurrency(amount int, stock string, target string) float64 {
	var result float64

	switch {
	case stock == "USD" && target == "EUR":
		result = float64(amount) / USDTOEUR
	case stock == "USD" && target == "RUB":
		result = float64(amount) / USDTORUB
	case stock == "EUR" && target == "RUB":
		result = float64(amount) / EURTORUB
	case stock == "EUR" && target == "USD":
		result = float64(amount) * USDTOEUR
	case stock == "RUB" && target == "USD":
		result = float64(amount) * USDTORUB
	case stock == "RUB" && target == "EUR":
		result = float64(amount) * EURTORUB
	default:
		result = 0
	}

	return result
}

func getStockCur() {
	for {
		fmt.Scan(&stockCur)
		if stockCur == "USD" || stockCur == "EUR" || stockCur == "RUB" {
			break
		} else {
			fmt.Print("Неизвестная валюта, попробуйте еще раз (USD, EUR, RUB): ")
		}
	}
}

func getAmount() {
	for {
		fmt.Scan(&amount)
		if amount >= 0 {
			break
		} else {
			fmt.Print("Число должно быть больше 0: ")
		}
	}
}

func getTargetCur() {
	for {
		fmt.Scan(&targetCur)
		if targetCur == "USD" || targetCur == "EUR" || targetCur == "RUB" {
			break
		} else {
			fmt.Print("Неизвестная валюта, попробуйте еще раз (USD, EUR, RUB): ")
		}
	}
}
