package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)


<<<<<<< HEAD

	go func() {
		ch <- 1
		b := <-ch
		fmt.Println("go func 里2",b)
=======
	go func() {
		ch <- 1
>>>>>>> 28f9a19cbe4a8af61d586b62e4170be6688c7635
	}()

	data := <-ch

	fmt.Println(data)

<<<<<<< HEAD
=======
	ch1 := make(chan int)

	go msg(ch1)
	fmt.Println("21行")
	for v := range ch1 {
		fmt.Println("ch",v)
	}
	fmt.Println("25行")

	//缓冲 管道
	ch3 := make(chan int)

	go msg(ch3)

	for v := range ch3 {
		fmt.Println("ch",v,"len")
	}

}



func msg (ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		ch <- i
		ch <- i
	}
	close(ch)
>>>>>>> 28f9a19cbe4a8af61d586b62e4170be6688c7635
}
