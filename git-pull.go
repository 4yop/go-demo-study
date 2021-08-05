package main

import (
	"fmt"
	"os/exec"
)

func main() {

	Command := []string{
		"git pull",
		"git add -A",
		"git commit-m go",
		"git push",
	}

	for _, s := range Command {
		c := exec.Command(s);
		err := c.Run()
		if (err != nil) {
			fmt.Println(err.Error())
		}
	}

	//


}
