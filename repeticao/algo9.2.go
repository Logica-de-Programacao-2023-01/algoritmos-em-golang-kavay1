package main

import "fmt"

func main() {

	var num, soma, contador float64
	fmt.Println("Digite uma sequência de números:")

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		soma += num
		contador++
	}

	fmt.Println("A média aritmética é: ", soma/contador)
	









}