package main

import "fmt"

func main()  {
	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
		case <-ch1:
			fmt.Println("<-ch1",<-ch1)
		case <-ch2:
			fmt.Println("<-ch2",<-ch2)
		case 
		default:
			fmt.Println("default")
	}
}