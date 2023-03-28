package main

import "fmt"

func main() {
	for num := 0; num <= 30; num++ {
		if num%3 == 0 {
			fmt.Println(num)
		}
	}
}
