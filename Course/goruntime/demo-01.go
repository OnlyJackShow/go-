package main

import (
	"fmt"
	"time"
)
//todo 并行-1
//并发执行的特点，是异步执行的
func hi()  {
	for i:=0;i<3;i++ {
		fmt.Printf("hello,golang开发者 %d\n",i)
	}
}

func jump(num int )  {
	for i:=0;i<num;i++ {
		fmt.Printf("hello,jump %d\n",i)
	}
}

func main() {

	go hi()
	go jump(3)
	fmt.Println("hello,world")

	time.Sleep(time.Second * 1) //它是等待中，可以把它理解为阻塞
}
