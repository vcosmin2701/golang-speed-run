package main

import "fmt"

func main(){

	const pi float64 = 3.141519
	var name string = "Ayyyyy"
	win := true
	x := 5

	fmt.Println(len(name))
	fmt.Println(name + " is a weird name.")

	fmt.Printf("%f \n", pi)
	fmt.Printf("%.3f \n", pi)

	fmt.Printf("%T \n", name)

	fmt.Printf("%t \n", win)
	fmt.Printf("%d \n", x)

	fmt.Printf("%b \n", 16)

	fmt.Printf("%c \n", 32) // ASCII

	fmt.Printf("%x \n", 15) // HEXA

	fmt.Printf("%e \n", pi)

}