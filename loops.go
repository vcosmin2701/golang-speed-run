package main

import "fmt"

func main(){

	// for loop
	for i := 1; i <= 10; i++{
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	// while loop
	i := 10
	for i<=20 {
		fmt.Printf("%d ", i)
		i++
	}

	// nested loops
	for i := 1; i < 5;i++{
		for j := 1; j < i; j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
}