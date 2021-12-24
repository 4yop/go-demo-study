package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var path string
func main() {
	path, _ := os.Getwd()
	fmt.Println("目录:",path)

	timer := time.NewTicker(180*time.Second)

	for {
		fmt.Println("开始咯",<-timer.C)

		output, err := runCmd("git", "add", "-A")
		if err != nil {
			fmt.Println("err1:", err)
		}
		fmt.Println(output)
		output, err = runCmd("git", "commit", "-m","go study")
		if err != nil {
			fmt.Println("err2:", err)
		}
		fmt.Println(output)
		output, err = runCmd("git", "push")
		if err != nil {
			fmt.Println("err3:", err)
		}
		fmt.Println(output)
	}
}

func runCmd(name string, arg ...string) (string,error) {
	 // 从配置文件中获取当前git仓库的路径

	cmd := exec.Command(name, arg...)
	cmd.Dir = path // 指定工作目录为git仓库目录
	//cmd.Stderr = os.Stderr
	msg, err := cmd.CombinedOutput() // 混合输出stdout+stderr
	cmd.Run()

	// 报错时 exit status 1
	return string(msg), err
}
