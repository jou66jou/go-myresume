package main

import (
	"fmt"
	"net"
	"runtime"
)

func main() {
	interfaceAddrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	for _, interfaceAddr := range interfaceAddrs {
		ipnet, ok := interfaceAddr.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				// var hostid string
				if runtime.GOOS == "windows" {
					// hostid = fmt.Sprintf(
					// 	"%.2x%.2x%.2x%.2x",
					// 	ipnet.IP[1],
					// 	ipnet.IP[0],
					// 	ipnet.IP[3],
					// 	ipnet.IP[2])
				} else {
					// hostid = fmt.Sprintf(
					// 	"%.2x%.2x%.2x%.2x",
					// 	ipnet.IP[13],
					// 	ipnet.IP[12],
					// 	ipnet.IP[15],
					// 	ipnet.IP[14])
				}
				fmt.Printf("ipnet : %v\n", ipnet.String())
			}
		}
	}
}
