package Service

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func SystemInfo() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	fmt.Println("主机名：", hostname)
	fmt.Println("系统类型：", runtime.GOOS)
	fmt.Println("系统架构：", runtime.GOARCH)
	fmt.Println("CPU 核心数：", runtime.GOMAXPROCS(0))
	netWork := NetWork()
	for _, data := range netWork {
		fmt.Println("网卡名称：", data.Name)
		fmt.Println("物理地址：", data.Mac)
		fmt.Println("网卡状态标识", data.Flags)
		for _, ipAddr := range data.IPAddrs {
			fmt.Println("IP", ipAddr.IP)
			fmt.Println("子网掩码", ipAddr.Mask)
		}
		println()
	}
}

func ParseMaskHexToMask(maskHex string) string {
	var mask string
	for index := range maskHex {
		if index%2 == 0 {
			hex := maskHex[index : index+2]
			num, _ := strconv.ParseInt(hex, 16, 0)
			mask = mask + strconv.FormatInt(num, 10)
		} else {
			if index < len(maskHex)-1 {
				mask = mask + "."
			}
		}
	}
	return mask
}

type InterfaceInfo struct {
	Name    string              `json:"name"`
	Mac     string              `json:"mac"`
	Flags   string              `json:"flag"`
	IPAddrs []InterfaceAddrInfo `json:"ipAddrs"`
}

type InterfaceAddrInfo struct {
	IP      string `json:"ip"`
	Mask    string `json:"mask"`
	MaskHex string `json:"maskHex"`
}

func NetWork() []InterfaceInfo {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	netWork := make([]InterfaceInfo, 0, 10)
	for _, _interface := range interfaces {
		interfaceAddrInfoSlice := make([]InterfaceAddrInfo, 0, 10)
		addrs, _ := _interface.Addrs()
		for _, addr := range addrs {
			interfaceAddrInfo := InterfaceAddrInfo{
				IP:      "0.0.0.0",
				Mask:    "0.0.0.0",
				MaskHex: "00000000",
			}
			if ip, ok := addr.(*net.IPNet); ok {
				if ip.IP.To4() != nil {
					interfaceAddrInfo.IP = ip.IP.String()
					maskHex := ip.Mask.String()
					interfaceAddrInfo.MaskHex = maskHex
					mask := ParseMaskHexToMask(maskHex)
					interfaceAddrInfo.Mask = mask
					interfaceAddrInfoSlice = append(interfaceAddrInfoSlice, interfaceAddrInfo)
				}
			}
		}

		interfaceInfo := InterfaceInfo{
			Name:    "UnKnown",
			IPAddrs: []InterfaceAddrInfo{},
			Mac:     "",
			Flags:   "",
		}
		interfaceInfo.Name = _interface.Name
		interfaceInfo.IPAddrs = interfaceAddrInfoSlice
		interfaceInfo.Mac = _interface.HardwareAddr.String()
		interfaceInfo.Flags = _interface.Flags.String()
		netWork = append(netWork, interfaceInfo)
	}
	return netWork
}

func TestCommand() {
	cmd := exec.Command("systeminfo")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Encode(out))
}
