package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main()  {
	wg.Add(2)

	go func1()
	go func2()

	fmt.Println("略略略略略略略")

	wg.Wait()

	fmt.Println("啊啊啊啊啊啊啊")

}

func func1 () {
	for i := 0; i < 10; i++ {
		fmt.Println("func1",i)
	}
	wg.Done()
}

func func2 () {
	for i := 0; i < 10; i++ {
		fmt.Println("func2",i)
	}
	wg.Done()
}