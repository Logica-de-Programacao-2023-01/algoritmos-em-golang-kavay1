package main

import "fmt"

func main() {

	var salario float64
	var hora float64

	fmt.Println("QUal seu salario mensal?")
	fmt.Scan(&salario)
	fmt.Println("qnts horas vc trabalha por mes?")
	fmt.Scan(&hora)

	fmt.Println("valor hora:", salario/hora)

}
