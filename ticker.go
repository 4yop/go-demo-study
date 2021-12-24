package main

import (
	"fmt"
	"time"
)

func main()  {
	//每隔x秒 执行向chan传时间
	timer := time.NewTicker(1*time.Second)
	a := <-timer.C
	fmt.Println(a)
	a = <-timer.C
	fmt.Println(a)
	a = <-timer.C
	fmt.Println(a)

	//timer.Stop()
	//timer.Reset(1*time.Second)

	time.Sleep(3*time.Second)
}