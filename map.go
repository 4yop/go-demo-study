package main

import (
	"fmt"
)

func main () {
	aaa := make(map[string]int)
	aaa["kkk"] = 1
	aaa["ppp"] = 2

	bbb := make(map[int]int);

	bbb[1] = 3
	bbb[10086] = 9
	bbb[2] = 5

	fmt.Println(aaa,bbb)
}
