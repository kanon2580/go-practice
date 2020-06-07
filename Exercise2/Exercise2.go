package main

import "fmt"
import "./temp"

func main() {
	fmt.Println(DibeCall())
}

func DibeCall() float64 {
	var beDivNum int = 20
	var divNum float64 = 0.002
	return temp.Dividing(beDivNum, divNum)
}
