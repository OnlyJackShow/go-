package main

import "fmt"

func main() {

	defer println("hello")

	defer println("world")

	defer println("test")

	fmt.Println("hello world")
}

