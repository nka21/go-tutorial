package main

import "fmt"

func val() {
	var i int = 10
	fmt.Println("i is", i)
	fmt.Printf("i type is %T, %v \n", i, i)

	var ui uint = 123
	fmt.Println("ui is", ui)
	fmt.Printf("ui type is %T, %v \n", ui, ui)

	// こっちの書き方の方が多い
	// 型推論を利用して、変数の型を自動的に決定している
	ii := 123 // int型
	// ii := uint(123)
	fmt.Printf("ii type is %T, %v \n", ii, ii)

	ff := 3.14 // float64型
	fmt.Printf("ff type is %T, %v \n", ff, ff)

	ss := "str" // string型
	fmt.Printf("ss type is %T, %v \n", ss, ss)
}