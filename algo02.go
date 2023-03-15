package main

import "fmt"

func main() {
	var num1 int

	fmt.Println("Digite um número")
	fmt.Scan(&num1)

	dobro := num1 * 2
	triplo := num1 * 3
	qradruplo := num1 * 4

	fmt.Println("O dobro do seu número é igual a:", dobro)
	fmt.Println("O triplo do seu número é igual a:", triplo)
	fmt.Println("O quadruplo do seu número é igual a:", qradruplo)

}
