package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("Qual seu salário?")
	fmt.Scan(&salario)

	aumento := 15 * salario / 100

	fmt.Println("Seu novo salário é: ", aumento+salario)

}
