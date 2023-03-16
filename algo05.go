package main

import "fmt"

func main() {
	var idade float64

	fmt.Println("Qual sua idade?")
	fmt.Scan(&idade)

	dias := 365 * idade
	fmt.Println("Sua idade em dias é:", dias)

}
