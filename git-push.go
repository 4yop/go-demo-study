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

	cmds := [][]string{
		{"git","pull"},
		{"git", "add", "-A"},
		{"git", "commit", "-m","go study"},
		{"git", "push"},
	}

	for {
		fmt.Println("-------------")
		fmt.Println("开始咯")
		for _, cmd := range cmds {
			fmt.Println(cmd)
			output, err := runCmd(cmd[0],cmd[1:]...)
			if err != nil {
				fmt.Println("err1:", err)
			}
			fmt.Println(output)
		}
		fmt.Println("-------------")
		time.Sleep(3*60*time.Second)
	}
}

func runCmd(name string, arg ...string) (string,error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = path // 指定工作目录为git仓库目录
	//cmd.Stderr = os.Stderr
	msg, err := cmd.CombinedOutput() // 混合输出stdout+stderr
	cmd.Run()

	// 报错时 exit status 1
	return string(msg), err
}
