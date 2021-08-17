package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)



	//ch <- 3


	go func() {
		a := <-ch
		fmt.Println("go func 里1",a)
		ch <- 1
		b := <-ch
		fmt.Println("go func 里2",b)
	}()

	data := <-ch

	fmt.Println(data)

}
