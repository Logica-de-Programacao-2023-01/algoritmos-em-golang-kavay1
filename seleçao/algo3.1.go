package main

import "fmt"

func main() {

	var num int

	fmt.Println("Digite um numero:")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("Seu numero é par:")
	} else {
		fmt.Println("Seu numero é impar:")
	}
}
