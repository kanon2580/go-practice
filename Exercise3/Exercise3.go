package main

import (
	"fmt"
)

func main() {
	var num int = 0
	for i := 1; i <= 100; i++ {
		num += i
	}
	fmt.Println(num)
}
