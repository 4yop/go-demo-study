package main

import "fmt"

func main()  {
	//defer fmt.Println("1")
	//defer fmt.Println("3")
	//fmt.Println("2")
	//defer fmt.Println("4")
	//
	//panic(errors.New("test"))

	for i:=0; i<5; i++{
		defer func(n int){
			fmt.Printf(string(n))
		}(i * 2)
	}
}
