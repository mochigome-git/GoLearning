package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	rect := Rectangle{width: 5.67, height: 3.44}
	fmt.Println("Area:", rect.area())
}
