package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	ch <- "from server1"
}
func server2(ch chan string) {
	ch <- "from server2"

}
func main() {
	output1 := make(chan string, 1)
	output2 := make(chan string, 10)
	/*go server1(output1)
	go server2(output2)*/
	/*	output1 <- "hello"
		output2 <- "world"*/
	output2 <- "world"
	time.Sleep(3 * time.Second)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	case <-time.After(1 * time.Second):
		goto timeout
	}

timeout:
	fmt.Println("超时了")

}
