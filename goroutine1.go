package main

import (
	"fmt"
	"sync"

)


var cat = make(chan int)
var eat = make(chan int)
var fish = make(chan int)

func main()  {
	cat<-0
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go printCat(&wg1)
	wg1.Add(1)
	go printEat(&wg1)
	wg1.Add(1)
	go printFish(&wg1)

	wg1.Wait()
	//for  {
	//	if _,ok := <-fish;ok == false {
	//		break;
	//	}
	//}
	defer close(cat)
	defer close(eat)
	defer close(fish)
}


func printCat(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		<-cat
		fmt.Printf("猫:%d\n",i)
		eat<-i
	}
	wg.Done()
}

func printEat(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		<-eat
		fmt.Printf("吃:%d\n",i)
		fish<-i
	}
	wg.Done()
}

func printFish(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		<-fish
		fmt.Printf("鱼:%d\n",i)
		cat<-i
	}
	wg.Done()

}