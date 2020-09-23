package main

import (
	"bytes"
	"strings"
)

func main() {
	
}

// go test -bench=. -benchmem
//go test -bench=.StringBuilder  -benchmem

// 网上的解释，很操蛋，所以，不能一概而论相信百度或谷歌出来的东西，
//对比  https://studygolang.com/articles/18764?fr=sidebar
//网上的测试数据，都是以万为单位,但是在实际的开发中，很少用到这么大量的字符串拼接

const count = 4

//加号
func StringJia(a string)string  {
	b:=""
	for i:=1;i<count;i++ {
		b+=a
	}
	return b
}

//StringBuilder
func StringBuilder(a string)string  {
	var b strings.Builder
	for i:=1;i<count;i++ {
		b.WriteString(a)
	}
	return b.String()
}

//bytes
func ByteString(a string)string {
	var b bytes.Buffer
	for i:=1;i<count;i++ {
		b.WriteString(a)
	}
	return b.String()
}


//todo 通过不断的调试，对结果进行分析,得出如下结论
// 1、第一个临界点是5次，意思是指，做字符串拼接的时候，直接用“+”号来做拼接，如果操作5次以内的次数（不包含5次），性能优于 strings.Builder和bytes.Buffer，达到5次后strings.Builder性能是最佳的，比其他的两个都好
// 2、第二个临界点是15次，意思是指，做字符串拼接的时候，直接用“+”号来做拼接，如果操作15次以内的次数（不包含15次），性能优于 strings.Builder和bytes.Buffer，达到15次后bytes.Builder性能是最佳的，比其他的两个都好

