
package main

import (
	"fmt"
	"reflect"
)

type City string
type Num int

func main()  {
	c :=  City("上海")
	n := Num(100)
	type1 := reflect.TypeOf(c);

	fmt.Println(c,n,type1)



	var cb = "123"
	fmt.Println(reflect.TypeOf(cb))
}