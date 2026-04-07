package main

import(
	"fmt"
)

func main(){
	const USDT = 100
    
    var EUR = 0.8643
	convert1 := float64(USDT) * EUR
    fmt.Println(convert1)
    
	var RUB =  78.35
	convert2 := float64(USDT) * RUB
	fmt.Println(convert2)

	convert3 := convert2 / convert1
	fmt.Println(convert3)
}
