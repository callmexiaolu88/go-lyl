package main

import (
	"fmt"
	"net"
)

func main() {
	ips := IPs()
	if ips != nil {
		for _, v := range ips {
			fmt.Println(v)
		}
	}
	lis, err := net.Listen("tcp", ips[0]+":0")
	if err != nil {
		fmt.Println(err)
	} else {
		addr := lis.Addr()
		if t, ok := addr.(*net.TCPAddr); ok {
			if t.IP.To4() != nil {
				fmt.Println(t.IP.To4())
			}
		}
	}
}

func avalidIP4() (ips []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, add := range addrs {
		if ipnet, ok := add.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return
}
