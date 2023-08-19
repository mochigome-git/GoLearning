package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
