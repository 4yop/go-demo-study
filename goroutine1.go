package main

import (
	"fmt"
	"time"
)


var cat = make(chan int)
var eat = make(chan int)
var fish = make(chan int)

func main()  {


	go printCat()
	go printEat()
	go printFish()
	cat<-0

	//for  {
	//	if _,ok := <-fish;ok == false {
	//		break;
	//	}
	//}

	time.Sleep(2*time.Second)
}


func printCat() {
	for i := 0; i < 10; i++ {
		<-cat
		fmt.Printf("猫:%d\n",i)
		eat<-i
	}
}

func printEat() {
	for i := 0; i < 10; i++ {
		<-eat
		fmt.Printf("吃:%d\n",i)
		fish<-i
	}
}

func printFish() {
	for i := 0; i < 10; i++ {
		<-fish
		fmt.Printf("鱼:%d\n",i)
		cat<-i
	}
	close(fish)
}