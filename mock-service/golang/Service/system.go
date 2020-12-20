package Service

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
)

func SystemInfo() {
	NetWork()
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	fmt.Println("主机名：", hostname)
	fmt.Println("系统类型：", runtime.GOOS)
	fmt.Println("系统架构：", runtime.GOARCH)
	fmt.Println("CPU 核心数：", runtime.GOMAXPROCS(0))
}

func NetWork() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(nil)
	}
	for _, _interface := range interfaces {
		fmt.Println(_interface)
	}
}

func TestCommand() {
	cmd := exec.Command("systeminfo")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Encode(out))
}
