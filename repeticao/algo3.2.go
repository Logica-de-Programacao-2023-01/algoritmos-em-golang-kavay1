package main

import "fmt"

func main() {

	for num := 1; num <= 19; num++ {
		if num%2 == 1 {
			fmt.Println(num)
		}
	}
}