
package main

import (
	"fmt";

)

type ListNode struct {
	val int
	next *ListNode

}

type People1 struct {
	Age int
	Name string
	Sex string
}


func main(){
	golang := People1{Age: 18,Name: "golang",Sex: "man"}

	xiaoming := People1{100,"xiaoming","woman"}





	fmt.Println(golang,xiaoming)

	//arr := [5]int{1,2,3,4,5}

}

type arr [5] int;

