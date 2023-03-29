package main

import "fmt"

func main() {

	for num := 1; num <= 100; num++ {
		if num%5 == 0 && num%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if num%5 == 0 {
			fmt.Println("Buzz")
		} else if num%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(num)
		}
	}

}
