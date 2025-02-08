package main

import "fmt"

func HelloWorld() {
	fmt.Println("hello world!")
}

func main() {
	HelloWorld()

	fmt.Println(sum(1, 2))
	fmt.Println(sub(5, 1))
}
