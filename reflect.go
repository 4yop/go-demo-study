package main

import (
	"fmt"
	"reflect"
)

type people struct {
	name string
	age int
}

func main()  {
	var a float64 = 123
	value := reflect.ValueOf(a)
	fmt.Println(value)
	i3 := value.Interface().(float64)

	abc := people{}
	fmt.Println(reflect.TypeOf(abc))

	fmt.Println(i3)
}