package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmdAdd()
}

func cmdAdd() {
	u := fmt.Sprintf("http://127.0.0.1/proxy.js")

	err := exec.Command("reg", "add", `HKCU\Software\Microsoft\Windows\CurrentVersion\Internet Settings`, "/v", "AutoConfigURL", "/t", "REG_SZ", "/d", u, "/f").Run()
	if err != nil {
		fmt.Printf("set AutoConfigURL error, %s\n", err.Error())
		return
	}
}
