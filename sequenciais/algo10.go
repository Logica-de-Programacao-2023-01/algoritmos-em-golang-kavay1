package main

import "fmt"

func main() {

	var peso float64

	fmt.Println("Qual seu peso?")
	fmt.Scan(&peso)

	fmt.Println("Seu peso em libras é de:", peso*2.2046)

}
