package main

import "fmt"

func main() {
	// stdin := bufio.NewReader(os.Stdin)

	// var a int
	// var b int

	// n, err := fmt.Scanf("%d %d", &a, &b)
	// if err != nil { // 에러가 발생하면
	// 	fmt.Println(n, err)
	// 	stdin.ReadString('\n')
	// } else { // 정상적으로 입력받으면
	// 	fmt.Println(n, a, b)
	// }

	// fmt.Print(":)")

	var a = 345
	var b = 3.1415

	fmt.Printf("%05d\n", a)
	fmt.Printf("%5.2f\n", b)
	fmt.Printf("%05.2f", b)
}
