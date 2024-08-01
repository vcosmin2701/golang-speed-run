package main

import "fmt"

func main(){
	var a int = 4
	var b float32 = 4.32
	const pi float64 = 3.14151394

	x,y := 14,15

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(pi)
	fmt.Println(x)
	fmt.Println(y)

	q,r := 5, 6

	fmt.Println(q + r);

	isbool := true
	internship := false;

	fmt.Println(isbool && internship);
	fmt.Println(!internship);
}