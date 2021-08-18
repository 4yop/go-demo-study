package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var a float64 = 123
	value := reflect.ValueOf(a)
	fmt.Println(value)
	i3 := value.Interface().(float64)

	fmt.Println(i3)
}