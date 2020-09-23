package main

import "fmt"

func main() {
	deferFuncParameter()
}

//todo 延迟函数的参数在defer语句出现的时候就已经确定了，所以无论后面如何修改alnt变量都不会影响延迟函数。
func deferFuncParameter() {
	var aInt = 1

	defer fmt.Println(aInt)

	aInt = 2
	return
}