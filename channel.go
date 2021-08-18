package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	fmt.Println(ch)


	go func() {
		ch <- 1
	}()

	fmt.Println(<-ch)

	ch1 := make(chan int)

	go msg(ch1)
	fmt.Println()
	for v := range ch1 {
		fmt.Println("ch",v)
	}



}


func msg (ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
