package main

import (
	"fmt"
	"time"
)

func main()  {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3*time.Second)
		ch1 <- 99

	}()

	select {
		case <-ch1:
			fmt.Println("<-ch1")
		case <-ch2:
			fmt.Println("<-ch2")
		case <-time.After(3*time.Second):
			fmt.Println("3333")
		//default:
		//	fmt.Println("default")
	}
}