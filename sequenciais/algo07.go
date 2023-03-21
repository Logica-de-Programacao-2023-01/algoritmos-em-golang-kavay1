package main

import "fmt"

func main() {

	var num int

	fmt.Println("Digite um número:")
	fmt.Scan(&num)

	adicao := num + 1
	sub := num - 1

	fmt.Println("Resultado:", adicao, sub)

}
