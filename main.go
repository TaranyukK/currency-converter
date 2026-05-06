package main

import "fmt"

var (
	amount              int
	stockCur            string
	targetCur           string
	convertationAmounts = map[string]map[string]float64{
		"USD": {
			"EUR": 0.86,
			"RUB": 75.5,
		},
		"EUR": {
			"USD": 1.17,
			"RUB": 88.29,
		},
		"RUB": {
			"USD": 0.013,
			"EUR": 0.011,
		},
	}
)

func main() {
	amount, stockCur, targetCur := getUserInput()

	resultNum := convertCurrency(amount, stockCur, targetCur, &convertationAmounts)

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

func convertCurrency(amount int, stock string, target string, rates *map[string]map[string]float64) float64 {
	var result float64

	result = float64(amount) * (*rates)[stock][target]

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
