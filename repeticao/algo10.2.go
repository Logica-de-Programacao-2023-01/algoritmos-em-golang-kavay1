package main

import "fmt"

func main() {

	var num int
	max := 0

	fmt.Println("Digite uma sequencia de números: ")

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		if num > max {
			max = num
		}

	}

	fmt.Println("O maior número é: ", max)

}
