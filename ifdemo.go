package main

import "fmt"

func main()  {
	var flag int
	fmt.Println("输入值:")
	fmt.Scanln(&flag)


	if flag == 3 {
		fmt.Println("3")
	} else if flag == 2 {
		fmt.Println("2")
	} else {
		fmt.Println("不2不3")
	}

	//仅对{}内 有 n
	if n := 4;n > 1{
		fmt.Println(n)
	}

}