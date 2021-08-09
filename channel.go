package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	fmt.Println(ch)

	ch1 := make(chan int)
	fmt.Println(ch1)
	go func() {
		ch <- 1
		//value := <-ch
		ch1 <- 1
	}()

	fmt.Println(<-ch)

	fmt.Println(<-ch1)
}
