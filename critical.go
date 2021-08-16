package main

import (
	"fmt"
	"time"
)

func main () {

	val := "临界值"

	go func() {
		val := "23131"
		fmt.Println(val)
	}()
	val = "1dfdf"
	time.Sleep(2 * time.Second)
	fmt.Println(val)
}