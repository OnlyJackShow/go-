package main

import "fmt"

func main() {

	//申明切片

	//方式1
	var ar1 []int
	ar1 = append(ar1, 1)
	fmt.Printf("%T \n", ar1)
	//方式2
	var ar2 = []int{1, 2, 3}
	fmt.Printf("%T---%v\n", ar2, ar2)
	//方式3
	var ar3 = []int{1: 1, 3: 2, 2: 3}
	fmt.Printf("%T---%v\n", ar3, ar3)
	//方式4
	ar4 := make([]int, 0)
	fmt.Printf("%T---%v\n", ar4, ar4)



}
