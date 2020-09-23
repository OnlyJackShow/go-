package main

import (
	"fmt"
	"time"
)

//todo 并发-01

func main() {
	go printer("hello")
	go printer("world")
	for{
		;
	}
}


func printer(str string){
	for _,val := range str{
		fmt.Printf("%c", val)
		time.Sleep(time.Millisecond * 300)
	}
}