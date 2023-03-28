package main

import "fmt"

func main() {

	for num := 0; num <= 20; num++ {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}

}