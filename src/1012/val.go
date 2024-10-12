package main

import "fmt"

func main() {
	var i int = 10
	fmt.Println("i is", i)
	fmt.Printf("i type is %T, %v \n", i, i)

	var ui uint = 123
	fmt.Println("ui is", ui)
	fmt.Printf("ui type is %T, %v \n", ui, ui)
}