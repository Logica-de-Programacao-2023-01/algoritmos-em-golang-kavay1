package main

import "fmt"

func main() {

	var dias float64
	var diaria float64

	fmt.Println("Quantos dias você trabalha?")
	fmt.Scan(&dias)

	fmt.Println("Qual valor de sua diaria? R$")
	fmt.Scan(&diaria)

	fmt.Println("Seu salario é: R$", dias*diaria)

}
