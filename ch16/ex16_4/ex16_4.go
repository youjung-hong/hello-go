package ex16_4

import (
	"fmt"
)

func changeArray(array2 [5]int) {
	array2[2] = 200
	printArray("변경 중", array2)
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
	printSlice("변경 중", slice2)
}

func Ex16_4() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	printArray("변경 전", array)
	printSlice("변경 전", slice)

	changeArray(array)
	changeSlice(slice)

	printArray("변경 후", array)
	printSlice("변경 후", slice)
}

func printArray(name string, array3 [5]int) {
	fmt.Printf("[%s] array=%v &array=%p &array[0]=%p\n", name, array3, &array3, &array3[0])
}

func printSlice(name string, slice3 []int) {
	fmt.Printf("[%s] slice=%v &slice=%p &slice[0]=%p\n", name, slice3, &slice3, &slice3[0])
}

func Ex16_7() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:3:4]

	printSlice("slice1", slice1)
	printSlice("slice2", slice2)
}
