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

	var c int = b // int에 float64 할당할 수 없다
	d := a * b    // int * float64 할 수 없다

	var e int64 = 7
	f := a * e // int * int64 할 수 없다

	fmt.Print(a, b, c, d, e, f)

	// ./main.go:15:14: cannot use b (variable of type float64) as int value in variable declaration
	// ./main.go:16:7: invalid operation: a * b (mismatched types int and float64)
	// ./main.go:19:7: invalid operation: a * e (mismatched types int and int64) //
}
