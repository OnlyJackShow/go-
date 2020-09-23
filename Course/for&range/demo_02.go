package main

import "fmt"

type Test struct {
	Index int
	Num   int
}

func main() {
	var t []Test
	t = append(t, Test{Index: 1, Num: 1})
	t = append(t, Test{Index: 2, Num: 2})
	for _, v := range t {
		v.Num = 100
		fmt.Println(v.Index, v.Num)
	}

	fmt.Println(t)
	for _, v := range t {
		fmt.Println(v.Index, v.Num)
	}


	for	i:= 0;i< len(t);i++{
		t[i].Num=100+i
	}
	fmt.Println(t)


/*	for i:= 0;i< len(t);i++{
		t[i].Num+=100
	}

	for _, v := range t {
		fmt.Println(v.Index, v.Num)
	}*/

/*	for i := range t {
		fmt.Println(t[i])
	}*/
}
