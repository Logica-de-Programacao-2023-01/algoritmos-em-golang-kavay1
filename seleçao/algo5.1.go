package main

import "fmt"

func main() {

	var num int

	fmt.Println("Digite um número inteiro:")
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("O número", num, "é multiplo de 5 e 3 ao mesmo tempo")

	} else {
		fmt.Println("O número", num, "não é multiplo de 5 e 3 ao mesmo tempo")
	}

}
