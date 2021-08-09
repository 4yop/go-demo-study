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


	//几种赋值方式
	var num1 int;
	num1 = 1

	var num2 int = 2

	var num3 = 3

	num4 := 4

	var num5,num6,num7 int
	num5 = 5
	num6 = 6
	num7 = 7

	var num8,num9 int = 8,9

	var num10,num11,num12 = 13,1.1,"qq"

	var (
		num13 = "num13"
		num14 = 14
		num15 int = 15
	)
	num16,num17 := "num16",17

	fmt.Println(num1,num2,num3,num4)
	fmt.Println(num5,num6,num7,num8,num9)
	fmt.Println(num10,num11,num12)
	fmt.Println(num13,num14,num15)
	fmt.Println(num16,num17)
	fmt.Println("num15的内存地址:",&num15)
}


