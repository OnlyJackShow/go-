package main

import "fmt"

//todo 无缓冲的

//todo 不仅仅是向 c 通道放 100，而是一直要等有别的携程 <-c1 接手了这个参数，那么c1<-100才会继续下去，要不然就一直阻塞着。

func main() {
	c := make(chan int)
	x := 100

	//c<-x

	go func() {
		c <- x
	}()
	fmt.Println(<-c)
}
