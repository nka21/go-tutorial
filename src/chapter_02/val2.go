package main

import "fmt"

func val2() {
	i := 1
	// fmt.Println("i is", i)
	fmt.Printf("i is %v, %p \n", i, &i)
	{
		i = 12
		fmt.Printf("in block i is %v, %p \n", i, &i)

		i := 123
		fmt.Printf("in block i is %v, %p \n", i, &i)

		ii := 123
		fmt.Println("in block ii is", ii)
	}
	// fmt.Println("in block ii is", ii) // undefined
	// fmt.Println("exit i is", i)
	fmt.Printf("exit i is %v, %p \n", i, &i)
}