package ex16_7

import (
	"fmt"
)

func printSlice(name string, slice3 []int) {
	fmt.Printf("[%s] slice=%v &slice=%p &slice[0]=%p\n", name, slice3, &slice3, &slice3[0])
}

func Ex16_7() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:3:4]

	printSlice("slice1", slice1)
	printSlice("slice2", slice2)

	fmt.Println("slice2[0]: ", slice2[0], "cap: ", cap(slice2))
	fmt.Println("slice2[1]: ", slice2[1], "cap: ", cap(slice2))

	slice3 := slice2[:3]
	slice4 := slice2[:4] // panic - cap까지만 slice할 수 있다
	fmt.Println("slice2", slice2, "slice3", slice3, "slice4", slice4)

	printSlice("slice1", slice1)
	printSlice("slice2", slice2)
	printSlice("slice3", slice3)
	printSlice("slice4", slice4)
}
