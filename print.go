package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//fmt.Println(enum.SHOW);
	//res := enum.GetStatus(100);
	//fmt.Println(res)

	var (
		a int;
		b float64
	);
	fmt.Println("请输入a和b:");
	fmt.Scanln(&a,&b);

	fmt.Println(a,b)

	fmt.Println("输入字符串:")

	reader := bufio.NewReader(os.Stdin)
	str,err :=reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ;
	}
	fmt.Println(str)
}

