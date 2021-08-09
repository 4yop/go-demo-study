package main

import "fmt"

type T struct{
	name string
}
func (t *T) F()  {
	t.name = "123"
}

func main()  {
	a, _ := test()
	fmt.Println(a)

	fmt.Println(test1())
	t := T{}
	t.F()

	fmt.Println(t)
}



func test () (int,string){
	return 1,"lzh"
}

func test1() (age int,name string){
	age =1
	name = "sdsdf"
	return
}

