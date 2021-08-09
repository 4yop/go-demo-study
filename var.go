package main

import (
	"fmt";
)

func main() {

	var a1 int = 2
	var a2 int = 3

	fmt.Println(a1/a2)

	var f1 float32 = 2
	var f2 float32 = 3

	fmt.Println(f1/f2)

	var (
		a,b int
		c float32
	)

	a = int(1)

	fmt.Println(a,b,c)
	fmt.Println("a")

	var arr[3][2] int;

	fmt.Println(arr)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice: ", slice)
	var start = 0
	var end = 3
	fmt.Println("第0个开始,3结束：",slice[start:end])

	slice1 := []int{1, 2, 3, 4, 5}

	fmt.Println(slice1[1:4])
}


