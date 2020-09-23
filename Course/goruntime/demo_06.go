package main

import (
	"fmt"
	"time"
)

// todo chan实现
func main() {
	var ch = make(chan string)
	str:="hello,world"
	for i:=0;i< len(str);i++ {
		go chanprint(string(str[i]), ch)
		m := <-ch
		fmt.Print(m)
	}
}

func chanprint(para string,ch chan  string) {
	time.Sleep(time.Millisecond*300)
	ch <- para
}