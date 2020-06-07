package main

import (
	"fmt"
)

func main() {
	var result []int
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			result = append(result, i)
		}
	}
	fmt.Println(result)
}
