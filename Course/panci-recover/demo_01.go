package main
import "fmt"

//todo 我们先看效果
func main() {
	defer func() {
		if info := recover(); info != nil {
			fmt.Println("触发了宕机", info)
		} else {
			fmt.Println("程序正常退出")
		}
	}()
	fmt.Println("bbbbbb")
	fmt.Println("cccccc")
	panic("fatal error")
	fmt.Println("ddddd")
	defer func() {
		fmt.Println("eeeeeeee")
	}()
	for {
		;
	}
}