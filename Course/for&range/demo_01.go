package main

import "fmt"

func main() {
	var arr_01 = [5]int{11, 22, 33, 44, 55}
	for _, v := range arr_01 {
		v = v + 1
	}
	fmt.Println(arr_01)
	for i := 0; i < len(arr_01); i++ {
		arr_01[i] = arr_01[i] + 1
	}
	fmt.Println(arr_01)
}
