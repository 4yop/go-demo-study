package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)



	ch <- 3


	go func() {
		fmt.Println("go func 里1",<-ch)
		ch <- 1
		fmt.Println("go func 里2",<-ch)
	}()

	data := <-ch

	fmt.Println(data)

}
