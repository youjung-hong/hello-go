package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 2, 3}
	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Println()
	fmt.Println(len(arr))
}
