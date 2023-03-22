package main

import "fmt"

func main() {

	var n1, n2 int

	fmt.Println("Digite dois números:")
	fmt.Scan(&n1, &n2)

	if n1 >= 0 && n2 >= 0 {
		fmt.Println("O resultado da multiplicação é:", n1*n2)
	} else {
		fmt.Println("O resultado da soma é:", n1+n2)
	}

}
