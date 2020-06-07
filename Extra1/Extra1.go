package main

import (
	"fmt"
)

func main() {
	var multiple_of_three []string
	var multiple_of_five []float64
	var multiple_of_fifteen []int

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			var element = fmt.Sprintf("%d個", i)
			multiple_of_three = append(multiple_of_three, element)
		}
		if i%5 == 0 {
			var element = float64(i) / 10
			multiple_of_five = append(multiple_of_five, element)
		}
		if i%15 == 0 {
			multiple_of_fifteen = append([]int{i}, multiple_of_fifteen...)
		}
	}

	fmt.Println(multiple_of_three)
	fmt.Println(multiple_of_five)
	fmt.Println(multiple_of_fifteen)
}
