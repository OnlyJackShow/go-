package main

import "fmt"

func main() {
	deferFuncParameter_01()
}

func printArray(array *[3]int) {
	for i := range array {
		fmt.Println(array[i])
	}
}

//todo 延迟函数的参数在defer语句出现时就已经确定了，即数组的地址，由于延迟函数执行时机是在return语句之前，所以对数组的最终修改值会被打印出来。
func deferFuncParameter_01() {
	var aArray = [3]int{1, 2, 3}
	defer printArray(&aArray)
	aArray[0] = 10
	return
}
