package main

import "fmt"

func main() {
	// var a int = 3
	// var b int
	// var c = 4
	// d := 5

	// fmt.Print(a, b, c, d)
	a := 3
	var b float64 = 3.5

	var c int = int(b)  // float64를 int로 형변환
	d := float64(a * c) // int * int를 float64로 형변환

	var e int64 = 7
	f := int64(a) * e // int64 * int64

	fmt.Print(a, b, c, d, e, f)
}
