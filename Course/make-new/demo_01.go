package main

import (
	"fmt"
	"reflect"
	"time"
)

//todo 用make申明，返回的是一个类型，而用new申明返回的是一个地址

//map slice chan 使用make创建最佳


func MakeFunc() {
	var s []string
	s = append(s, "aa")
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d, %v\n", len(s), cap(s), s)
	fmt.Println("type=", reflect.TypeOf(s))
}

//使用map必须使用make
//map 是 引用类型 的： 内存用 make 方法来分配
func MakeMap() {
	//todo 这个代码肯定会报错，为什么会报错，是因为你只是申明了它，没有给它分配内存，然后直接给它赋值了
	/*var m map[string]int
	m["aa"] = 11
	fmt.Println(m)*/

	/*//改进
	var m map[string]int
	m=make(map[string]int)
	m["aa"] = 11
	fmt.Println(m)*/
}


func MakeChan()  {
	//这里channel要和goruntine结合起来使用更好
	p:=make(chan int)
	go func() {
		p<-111
	}()
	mm:=<-p
	fmt.Println("mm=",mm)
	var op chan string
	op<-"111"
	fmt.Println("op==",<-op)
}


func MakeCap()  {
	var mmap=make(map[string]string ,1)
	mmap["a"]="1"
	mmap["b"]="2"
	mmap["c"]="3"
	mmap["d"]="4"
	mmap["e"]="5"
	mmap["f"]="6"
	mmap["g"]="7"
	mmap["h"]="8"
	mmap["i"]="9"
	mmap["j"]="10"
	mmap["k"]="11"
	mmap["l"]="12"
	mmap["m"]="13"
	fmt.Println("mmap=",mmap)
}

func Testtest() {
	var serverDone = make(chan string)

	var ticker = time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	var counter = 0
	for {
		select {
		case <-serverDone:
			return
		case <-ticker.C:
			counter += 1
			if counter == 10 {
				goto Exit
			}
		}
	}
Exit:
	fmt.Println("结束了")
}