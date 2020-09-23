package main

import "fmt"

//todo 有缓冲的

//todo c2<-99 则不会阻塞，因为缓冲大小是1(其实是缓冲大小为0)，只有当放第二个值的时候，第一个还没被人拿走，这时候才会阻塞。
func main() {
	c2:=make(chan int,1)
	c2<-99
	//c2<-100
	fmt.Println(<-c2)
}
