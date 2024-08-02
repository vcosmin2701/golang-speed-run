package main

import "fmt"

func main(){

	age := 18
	if age >= 18 {
		fmt.Println("You are able to drive")
	}else{
		fmt.Println("Grow the heck up!")
	}

	switch age {
	case 16: fmt.Println("B1 is trash")
	default: fmt.Println("Do you have a driving license ?")
	}
}