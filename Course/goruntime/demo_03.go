package main

import (
	"fmt"
	"sync"
)

//todo 并行-3

var wgg sync.WaitGroup

func main() {
	wgg.Add(2)
	go bingxing1()
	go bingxing2()
	wgg.Wait()
}

//并行
func bingxing1()  {
	defer 	wgg.Done()
	fmt.Println("我是并行1")
}

//并行
func bingxing2()  {
	defer 	wgg.Done()
	fmt.Println("我是并行2")
}