package main

import "fmt"

func main () {
	rect1 := Rectangle{10, 20}

	fmt.Println("Area of rectangle is: ", rect1.area())
}


type Rectangle struct {
	height float64
	width float64
}

func (rect *Rectangle) area() float64{
	return rect.height * rect.width
}