package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)
var name string;
func main()  {

	name = "小明"
	go printName()
	go printName()
	time.Sleep(time.Second)

	name = "小红"
	go printName()
	go printName()
	time.Sleep(time.Second)

	var (
		wg        sync.WaitGroup
		numbers[] int
		//mux       sync.Mutex
		rw        sync.RWMutex
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			rw.RLock()
			numbers = append(numbers,i)
			rw.RUnlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(numbers)

}

func printName()  {
	log.Println("名字是",name,"呀")
}