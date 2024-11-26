package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b)
	if err != nil { // 에러가 발생하면
		fmt.Println(n, err)
	} else { // 정상적으로 입력받으면
		fmt.Println(n, a, b)
	}
}
