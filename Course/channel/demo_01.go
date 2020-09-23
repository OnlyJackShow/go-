package main

import "fmt"


//todo val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值
func main() {
	params := "hello"
	ch := make(chan string)
	go func() {
		funcname(params,ch)
	}()
	/*	for s := range ch {
		fmt.Println("s=",s)
	}*/
	for i:=0;i< len(ch);i++ {
		fmt.Println("for0=",<-ch)
	}
	r:=<-ch
	fmt.Println("r=",r)
}

func funcname(str string,ch chan string) {
	if len(str) > 0 {
		ch <- str
	} else {
		ch <- "hello"
	}
}
