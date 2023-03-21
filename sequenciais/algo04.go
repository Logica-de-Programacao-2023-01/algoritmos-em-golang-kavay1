package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Println("Digite 3 números:")
	fmt.Scan(&num1, &num2, &num3)

	media := (num1*2 + num2*3 + num3*5) / 10
	fmt.Println("Sua media ponderada é:", media)

}
