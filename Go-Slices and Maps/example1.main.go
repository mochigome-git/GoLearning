package main

import "fmt"

func main() {
	// Slices
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", numbers)

	// Maps
	person := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	fmt.Println("Alice's Age:", person["Alice"])
}
