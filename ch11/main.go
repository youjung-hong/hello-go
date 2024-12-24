package main

import "fmt"

type Student struct {
	Age   int
	no    int
	Score float64
}

func main() {
	student := Student{15, 23, 88.2}

	student2 := student

	fmt.Println(student2)
}
