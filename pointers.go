package main

import "fmt"

func main () {
	x := 5
	// fmt.Println(&x) ->  the address in memory of x
	changeValue(&x)
	fmt.Println(x)

}

func changeValue(x *int){
	*x = 7;
}