package main

import "fmt"

func main() {
	// 配列
	// 長さを3から変えることが出来ない。
	x := [3]int{1,2,3}
	fmt.Println(x)
	fmt.Printf("x type is %T \n", x)
	// x = append(x, 5)

	// スライス
	// 長さを3から変えることが出来ない。
	y := []int{1,2,3}
	fmt.Println(y)
	fmt.Printf("y type is %T \n", y)
	y = append(y, 5)
	fmt.Println(y)

	// 
	z := make([]int, 0)
	fmt.Println(z)
	z = append(z, 11, 22, 33)
	fmt.Println(z)
}