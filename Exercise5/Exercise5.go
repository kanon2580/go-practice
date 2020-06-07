package main

import (
	"fmt"
)

func main() {
	var array1 [10]int
	for i := 0; i < 10; i++ {
		array1[i] = i + 1
	}
	fmt.Println(array1)
}
