package main

import "fmt"

func main()  {
	c := make(chan int,3)
	go func() {
		for i := 0; i < 3; i++ {
			c<-i
		}
		close(c)
	}()

	for i3 := range c {
		fmt.Println(i3)
	}

}