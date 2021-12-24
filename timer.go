package main

import (
	"fmt"
	"time"
)

func main()  {
	//睡2秒
	time.Sleep(2*time.Second)
	//创定时器，2秒后写入时间进通道
	timer := time.NewTimer(2*time.Second)
	t := <-timer.C
	fmt.Println(t)
	//重置定时器
	timer.Reset(1*time.Second)
	t = <-timer.C
	fmt.Println(t)
	//n := 1
	//timer1 := time.AfterFunc(4*time.Second, func() {
	//	n++
	//})
	////停止定时器
	//a := timer1.Stop()
	//
	//fmt.Println(a)
	//timer1.Stop()
}

