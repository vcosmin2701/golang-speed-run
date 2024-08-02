package main

import "fmt"

func main() {
	// var EvenNum[5] int
	EvenNum := [5]int{0,2,4,6,8}

	for _, value := range EvenNum {
		fmt.Printf("%d ", value);
	}

	numSlice := []int{5, 4, 3, 2, 1}
	sliced := numSlice[3:5]
	fmt.Println(sliced)

	slice2 := make([]int, 5, 10)
	fmt.Println(slice2)
	copy(slice2, numSlice)

	fmt.Printf("%T \n", numSlice)
	fmt.Printf("%T \n", sliced)
	fmt.Printf("%T \n", slice2)

	fmt.Println(numSlice)
	fmt.Println(sliced)
	fmt.Println(slice2)

	slice3 := append(numSlice, 3, -10)
	fmt.Println(slice3)

}