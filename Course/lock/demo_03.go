package main

import (
	"fmt"
	"sync"
	"time"
)

var mm *sync.RWMutex

func main() {
	mm = new(sync.RWMutex)

	// 写的时候啥也不能干
	go write(1)
	go readm(2)
	go write(3)

	for{
		;
	}
	//time.Sleep(2*time.Second)
}

func readm(i int) {


	println(i,"读开始")
	mm.RLock()
	println(i,"读取中")
	time.Sleep(1*time.Second)
	mm.RUnlock()

	fmt.Printf("goroutine %d 读结束，值为：%d\n", i, i)
}

func write(i int) {
	println(i,"写开始")

	mm.Lock()
	println(i,"写入中")
	time.Sleep(1*time.Second)
	mm.Unlock()

	println(i,"写结束")
}