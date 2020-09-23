package main

import (
	"fmt"
	"sync"
	"time"
)

var mut sync.Mutex

func printer(str string) {
	mut.Lock()
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Millisecond * 300)
	}
	mut.Unlock()
}

func main() {
	go printer("hello")
	go printer("world")
	for{
		;
	}

}