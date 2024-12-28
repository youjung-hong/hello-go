package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3)

	fmt.Println("slice1=", slice1)
	fmt.Println("slice2=", slice2)

	slice1[2] = 100
	fmt.Println("slice1=", slice1)
	slice2[2] = 100
	fmt.Println("slice2=", slice2)

	slice1[3] = 200
	fmt.Println("slice1=", slice1)
	slice2[3] = 200
	fmt.Println("slice2=", slice2)
}
