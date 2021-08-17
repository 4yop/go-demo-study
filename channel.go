package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)


	ch1 := make(chan int)



	go func() {
		ch <- 1
		//value := <-ch
		ch1 <- 1
		fmt.Println(<-ch)

	}()

	data := <-ch

	fmt.Println(data)

	fmt.Println(<-ch1)
}
