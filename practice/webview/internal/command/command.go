package command

import (
	"net"
	"runtime"
)

func GetIP() (ips []string, err error) {
	interfaceAddrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
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
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return ips, nil
}
