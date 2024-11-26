package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil { // 에러가 발생하면
		fmt.Println(n, err)
		stdin.ReadString('\n')
	} else { // 정상적으로 입력받으면
		fmt.Println(n, a, b)
	}

	fmt.Print(":)")
}
