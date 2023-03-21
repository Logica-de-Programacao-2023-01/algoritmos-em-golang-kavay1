package main

import "fmt"

func main() {

	var num1, num2 int

	fmt.Println("Digite dois numeros:")
	fmt.Scan(&num1, &num2)

	if num1 > num2 {
		fmt.Println("O maior numero é", num1)

	} else {
		fmt.Println("O maior numero é:", num2)
	}

}
