package main

import "fmt"

func main() {
	a :=  make([]int,0)
	n := 20
	fmt.Printf("len=%d cap=%d\n", len(a), cap(a))
	fmt.Println("//////////////////////////////////////////")
	for i := 0; i < n; i++ {
		a = append(a, 1)
		fmt.Printf("len=%d cap=%d\n", len(a), cap(a))
	}

}
