package main

import (
	"fmt"
	"time"
)

func main() {
	//声明一个channel
	ch := make(chan int)

	//声明一个匿名函数，传入一个参数整型channel类型ch
	go func(ch chan int) {
		ch <- 1
		//往channel写入一个数据，此时阻塞
	}(ch)

	//由于goroutine执行太快，先让它sleep 1秒
	time.Sleep(3 * time.Second)
	select {
	//读取ch，解除阻塞
	case <-ch:
		fmt.Println("come to read ch!")
	case <-time.After(1 * time.Second):
		goto timeout

	default:
		fmt.Print("come to default!")
	}
timeout:
	fmt.Println("超时了")
}
