package main

import (
	"fmt"
	"runtime"
)

func main()  {

	fmt.Println("go 安装在",runtime.GOROOT())

	//386、amd64 或 arm。
	fmt.Println("处理器架构",runtime.GOARCH)

	fmt.Println("系统",runtime.GOOS)





}
