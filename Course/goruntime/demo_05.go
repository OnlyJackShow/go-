package main

import (
	"fmt"
	"sync"
	"time"
)

//todo 并发-02
//  互斥锁实现
var look sync.Mutex

func main() {
	go printers("hello")
	go printers("world")
	for{
		;
	}
}

func printers(str string){
	//look.Lock()
	for _,val := range str{
		fmt.Printf("%c", val)
		time.Sleep(time.Millisecond * 300)
	}
	//look.Unlock()
}