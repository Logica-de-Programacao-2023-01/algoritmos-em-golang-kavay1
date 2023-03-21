package main

import "fmt"

func main() {

	var preço float64

	fmt.Println("Preço:")
	fmt.Scan(&preço)

	desconto := (preço * 10) / 100

	fmt.Println("O novo preço do produto é de:", preço-desconto)
}
