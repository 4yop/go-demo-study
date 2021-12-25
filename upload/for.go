package main

import "fmt"

func main()  {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}

	//for {
	//	fmt.Println("死循环")
	//}

	//终止 ： break out  continue out 层

	out:for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (j == 1) {
				continue out;
			}
			fmt.Println("i:%d,j:%d",i,j)
		}
	}

}