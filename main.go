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
}

func getUserInput() (int, string, string) {
	fmt.Print("Введите количество валюты, которую хотите ковертировать: ")
	fmt.Scan(&amount)
	fmt.Print("Введите валюту, которую хотите конвертировать: ")
	fmt.Scan(&stockCur)
	fmt.Print("Введите валюту, в которую хотите конвертировать: ")
	fmt.Scan(&targetCur)

	return amount, stockCur, targetCur
}

func convertCurrency(amount int, stock string, target string) {}
