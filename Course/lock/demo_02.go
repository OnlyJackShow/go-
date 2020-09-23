package main

import (
	"sync"
	"time"
)

//todo 读写锁之读

var m *sync.RWMutex


//todo 读的时候可以多个goruntime同时读
func main() {
	m = new(sync.RWMutex)

	// 多个同时读
	go read(1)
	go read(2)

	time.Sleep(2*time.Second)
}

func read(i int) {
	println(i,"read start")

	m.RLock()
	println(i,"reading")
	time.Sleep(1*time.Second)
	m.RUnlock()

	println(i,"read over")
}