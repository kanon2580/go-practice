package main

import "fmt"

func main() {
	var beDivNum int = 20
	var divNum float64 = 0.002
	fmt.Println(dividing(beDivNum, divNum))
}

func dividing(a int, b float64) float64 {
	return float64(a) / b
}
