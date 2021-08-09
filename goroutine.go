package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main()  {
	var wg  sync.WaitGroup
	wg.Add(3)
	go doSomeThing(1,&wg)
	go doSomeThing(2,&wg)
	go doSomeThing(3,&wg)

	wg.Wait()
	log.Printf("finish all jobs\n")


	fmt.Println("--------------------------------")
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)

	for i := 0; i < 3; i++ {
		go func(v int) {
			fmt.Println(v)
		}(i)
	}
	time.Sleep(1 * time.Second)
}

func doSomeThing(id int,wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("before do job:(%d) \n", id)
	time.Sleep(3 * time.Second)
	log.Printf("after do job:(%d) \n", id)
}