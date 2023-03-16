package main

import "fmt"

func main() {
	var peso float64
	var altura float64

	fmt.Println("Qual seu peso?")
	fmt.Scan(&peso)
	fmt.Println("Qual sua altura?")
	fmt.Scan(&altura)

	imc := peso / (altura * altura)
	fmt.Println("O seu IMC é:", imc)

}
