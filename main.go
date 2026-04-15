package main

import (
	"fmt"
)

func main() {

	USDT, EUR, RUB := getUserInput()

	var eurToUsdt float64 = 1.08
	var rubToUsdt float64 = 0.01

	convert1, convert2, convert3 := convertation(USDT, EUR, RUB, eurToUsdt, rubToUsdt )
	fmt.Printf("USDT к EUR: %.2f EUR\n", convert1)
	fmt.Printf("USDT к RUB: %.2f RUB\n", convert2)
	fmt.Printf("EUR к RUB: %.2f RUB\n", convert3)
}

func getUserInput() (float64, float64, float64) {
	var USDT, EUR, RUB float64

	fmt.Println("Введите количество USDT")
	fmt.Scan(&USDT)
	fmt.Println("Введите количество EUR")
	fmt.Scan(&EUR)
	fmt.Println("Введите количество RUB")
	fmt.Scan(&RUB)
	return USDT, EUR, RUB
}

func convertation(USDT, EUR, RUB, eurToUsdt, rubToUsdt float64) (float64, float64, float64) {
	convert1 := USDT / eurToUsdt
	convert2 := USDT / rubToUsdt
	eurInUsdt := EUR * eurToUsdt
	convert3 := eurInUsdt / rubToUsdt
	return convert1, convert2, convert3

}
