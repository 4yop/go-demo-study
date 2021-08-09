package main

import (
	"fmt"
	"go-demo-study/enum"
)


func main() {
	fmt.Println(enum.SHOW);
	res := enum.GetStatus(100);
	fmt.Println(res)
}

