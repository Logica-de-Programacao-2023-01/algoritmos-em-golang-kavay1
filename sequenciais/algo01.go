package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var num3 int

	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&num1)
	fmt.Println("Digite o segundo número:")
	fmt.Scan(&num2)
	fmt.Println("Digite o terceiro número:")
	fmt.Scan(&num3)

	resultado := num1 + num2 + num3
	fmt.Println("Seu resultado é:", resultado)

}
