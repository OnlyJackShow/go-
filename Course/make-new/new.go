package main

import (
	"fmt"
	"reflect"
)

func main() {
	a:=new(int)
	*a=11
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(*a)
}
