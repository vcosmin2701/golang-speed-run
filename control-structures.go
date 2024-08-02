package main

import "fmt"

func main () {
	// defer example, FirstRun will wait until
	
	// ###########

	// SecondRun() is executed
	// defer FirstRun()
	// SecondRun()


	// ###########

	// recover + panic 

	// fmt.Println(div(3,0))
	// fmt.Println(div(3,5))
	// demPanic()


	// function with multiple arguments
	fmt.Println(addThemUp(10, 20, 30, 40, 50))


}

func addThemUp (args ...int) int {
	sum := 0

	for _,value := range args {
		sum += value;
	}

	return sum
}

func div (num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1/num2
	return solution
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("Panic")
}

func FirstRun() { fmt.Println("I executed first")}

func SecondRun() { fmt.Println("I executed second")}