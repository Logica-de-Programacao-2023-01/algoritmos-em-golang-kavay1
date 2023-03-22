package main

import "fmt"

func main() {

	var salario float64

	fmt.Println("Qual valor do seu salário?")
	fmt.Scan(&salario)

	if salario < 1000 {
		fmt.Println("Novo salário com 10% de aumento é:", (salario/100)*10+salario)
	} else {
		fmt.Println("Novo salário com 5% de aumento é:", (salario/100)*5+salario)
	}

}
