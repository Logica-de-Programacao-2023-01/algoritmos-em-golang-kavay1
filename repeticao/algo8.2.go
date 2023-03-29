package main

import "fmt"

func main() {
	var num int
	fmt.Println("Digite um número:")
	fmt.Scan(&num)

	fmt.Println("Os divisores de", num, "são:")
	for div := 1; div <= num; div++ {
		if num%div == 0 {
		fmt.Print(div, " ")
			
		}

	}

}