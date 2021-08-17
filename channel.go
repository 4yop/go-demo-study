package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)



	go func() {
		ch <- 1
		b := <-ch
		fmt.Println("go func 里2",b)
	}()

	data := <-ch

	fmt.Println(data)

}
