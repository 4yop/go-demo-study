package main

import "fmt"

func main () {

	//延迟执行 defer ,先进后出
	defer fmt.Println("A")
	fmt.Println("C")
	defer fmt.Println("B")

	//for i := 0; i < 5; i++ {
	//	defer fmt.Println(i)
	//}
	for i := 0; i < 5; i++ {
		defer echo(i)
	}
}

func echo (i int) {
	fmt.Println(i)
}