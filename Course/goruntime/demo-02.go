package main

import (
	"fmt"
	"sync"
)
//todo 并行-2
//并发执行的特点，是

var wg sync.WaitGroup

func Hi()  {
	defer wg.Done()
	for i:=0;i<3;i++ {
		fmt.Printf("hello,golang开发者 %d\n",i)
	}
}

func jump1(num int )  {
	defer wg.Done()
	for i:=0;i<num;i++ {
		fmt.Printf("hello,jump %d\n",i)
	}
}

func main() {
	wg.Add(2)
	go Hi()
	go jump1(3)
	fmt.Println("hello,world")
	wg.Wait() //它是阻塞的
	//	time.Sleep(time.Second*1)
}
