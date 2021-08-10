package main

import (
	"fmt"
	"math/rand"
	"time"
)

var i =0
func main()  {

	var a = 10

	A:
		i++
	for i < a{
		fmt.Println(i)
		goto A
	}

	fmt.Println(rand.Int())

	fmt.Println("当前时间:",time.Now())

	fmt.Println("当前时间戳:",time.Now().Unix())

	fmt.Println("当前时间纳秒:",time.Now().UnixNano())



}