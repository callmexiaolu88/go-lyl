package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, _ := net.InterfaceAddrs()
	for _, add := range addrs {
		fmt.Println(add.String())
		switch t := add.(type) {
		case *net.IPNet:
			fmt.Printf("%s:\t%s\n", "IsLoopback", t.IP.IsLoopback())
			fmt.Printf("%s:\t%s\n", "IsLinkLocalMulticast", t.IP.IsLinkLocalMulticast())
			fmt.Printf("%s:\t%s\n", "IsLinkLocalUnicast", t.IP.IsLinkLocalUnicast())
			fmt.Printf("%s:\t%s\n", "IsInterfaceLocalMulticast", t.IP.IsInterfaceLocalMulticast())
			fmt.Printf("%s:\t%s\n", "IsMulticast", t.IP.IsMulticast())
			fmt.Printf("%s:\t%s\n", "IsGlobalUnicast", t.IP.IsGlobalUnicast())
			fmt.Printf("%s:\t%s\n", "IsUnspecified", t.IP.IsUnspecified())
		case *net.IPAddr:
			fmt.Printf("%s:\t%s\n", "IsLoopback", t.IP.IsLoopback())
			fmt.Printf("%s:\t%s\n", "IsLinkLocalMulticast", t.IP.IsLinkLocalMulticast())
			fmt.Printf("%s:\t%s\n", "IsLinkLocalUnicast", t.IP.IsLinkLocalUnicast())
			fmt.Printf("%s:\t%s\n", "IsInterfaceLocalMulticast", t.IP.IsInterfaceLocalMulticast())
			fmt.Printf("%s:\t%s\n", "IsMulticast", t.IP.IsMulticast())
			fmt.Printf("%s:\t%s\n", "IsGlobalUnicast", t.IP.IsGlobalUnicast())
			fmt.Printf("%s:\t%s\n", "IsUnspecified", t.IP.IsUnspecified())
		}
	}
}
