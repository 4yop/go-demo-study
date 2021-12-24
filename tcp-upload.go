package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("输入文件：")
	cmdReader := bufio.NewReader(os.Stdin)
	cmdStr, err := cmdReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	//这里把读取的数据后面的换行去掉，对于Mac是"\r"，Linux下面
	//是"\n"，Windows下面是"\r\n"，所以为了支持多平台，直接用
	//"\r\n"作为过滤字符
	cmdStr = strings.Trim(cmdStr, "\r\n")

	content,err := os.ReadFile(cmdStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
