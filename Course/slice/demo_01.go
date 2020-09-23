package main

import "fmt"

//todo 数组和切片

func main() {

	//数组 声明第一种方法
	var arr1 [3]int
	arr1[0]=1
	arr1[1]=2
	arr1[2]=3
	fmt.Println(arr1)

	//数组 声明第二种方法
	var arr2=[3]string{"11","22","33"}
	fmt.Println("arr2=",arr2)

	//数组 声明第三种方法
	var arr3=[...]string{"aa","bb","cc","dd"}
	fmt.Println("arr3=",arr3)

	//数组 声明第四种方法
	var arr4=[...]string{1:"aa",3:"bb",2:"cc",4:"dd"}
	fmt.Println("arr3=",arr4)
}

