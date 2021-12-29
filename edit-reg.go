package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os/exec"
)

func main() {

	key, exists, err := registry.CreateKey(registry.CURRENT_USER,  `Control Panel\Desktop`, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer key.Close()

	if exists {
		fmt.Println("键已存在")
	} else {
		fmt.Println("新建注册表键,自己去新建")
		return
	}

	// 读取字符串值 获取屏保时间
	s, _, _ := key.GetStringValue("ScreenSaveTimeOut")
	fmt.Println("当前ScreenSaveTimeOut的值为",s)

	newVal := "600"

	if s == newVal {
		fmt.Println("值一样不用改了")
		return
	}

	err = key.SetStringValue("ScreenSaveTimeOut", newVal)
	if err != nil {
		fmt.Println("修改失败",err)
		return;
	}

	fmt.Println("改成功为:",newVal,"等3秒重启")

	exec.Command("shutdown","-r","-t","3").Run()
}
