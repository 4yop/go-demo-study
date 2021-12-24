package main

import (
	"fmt"
	_ "github.com/CodyGuo/win"
	"golang.org/x/sys/windows/registry"
	"log"
	"os/exec"
	"time"
)

func main() {


	for {

		startTime := getTimestamp(9, 30)
		endTime := getTimestamp(18, 00)

		//t := time.Unix(startTime, 0).Format("2006-01-02 03:04:05 PM")
		//t1 := time.Unix(endTime, 0).Format("2006-01-02 03:04:05 PM")

		now := time.Now().Unix()

		if now >= startTime && now <= endTime {
			fmt.Println("上班了")
			editRegister(true)
		} else {
			fmt.Println("没上班")
			editRegister(false)
		}
		time.Sleep(60*time.Second)
		continue;
	}



}

//获取时间戳
func getTimestamp(hour int, min int) int64 {
	currentTime := time.Now()
	timestamp := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), hour, min, 0, 0, currentTime.Location()).Unix()
	return timestamp
}


//修改注册表并重启
//isWork 是否在上班，
func editRegister(isWork bool) {
	key, exists, err := registry.CreateKey(registry.CURRENT_USER,  `Control Panel\Desktop`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	if exists {
		fmt.Println("键已存在")
	} else {
		fmt.Println("新建注册表键")
		return
	}
	// 读取字符串值
	s, _, _ := key.GetStringValue("ScreenSaveTimeOut")
	fmt.Println("当前ScreenSaveTimeOut的值为",s)

	newVal := ""
	if !isWork {
		newVal = "60"
		fmt.Println("没在上班要改为",newVal)
	}else{
		newVal = "6000"
		fmt.Println("上班要改为",newVal)
	}

	if newVal == s {
		fmt.Println("旧值为:",s,",不用改")
		return;
	}

	err = key.SetStringValue("ScreenSaveTimeOut", newVal)
	if err != nil {
		fmt.Println("修改失败",err)
		return;
	}
	fmt.Println("改成功为:",newVal,"等2分钟重启")
	time.Sleep(120*time.Second)
	exec.Command("shutdown","-r","-t","3").Run()
}