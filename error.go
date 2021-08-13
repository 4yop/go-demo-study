package main

import (
	"errors"
	"fmt"
	"os"
)

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

	f,err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)

}

func checkAge(age int) error {
	if age > 0 {
		return nil
	}
	return errors.New("错误")
}
