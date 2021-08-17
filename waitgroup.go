package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Goods struct {
	stock int
}

func main()  {
	//wg.Add(10)
	//
	//go func1()
	//go func2()
	//
	//fmt.Println("略略略略略略略")
	//
	//wg.Wait()
	//
	//fmt.Println("啊啊啊啊啊啊啊")
	goods := Goods{stock: 10}
	fmt.Println(goods)

	for i := 0; i < 20; i++ {
		go goods.cutStock(i)
	}
	//goods.cutStock(2)
	time.Sleep(11*time.Second)
	fmt.Println(goods)
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


func (goods *Goods) cutStock (id int) {
	for i := 0; i < 12; i++ {
		if goods.stock <= 0 {
			fmt.Println("id"+strconv.Itoa(id)+" cutStock :","没有库存了")
			break;
		}
		time.Sleep(1*time.Second)
		goods.stock--;
		stock := goods.stock
		fmt.Println("id"+strconv.Itoa(id)+" cutStock 减1:",stock)
	}
	//defer wg.Done()
}