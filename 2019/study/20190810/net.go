package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.JoinHostPort("0.0.0.0", "8888"))
	fmt.Println(net.SplitHostPort("127.0.0.1:9999"))
	fmt.Println(net.LookupAddr("180.101.49.12"))
	fmt.Println(net.LookupHost("www.baidu.com"))
	fmt.Println(net.ParseCIDR("192.168.1.1/24"))

	ip := net.ParseIP("127.0.0.1")
	fmt.Printf("%T, %v\n", ip, ip)

	ips, err := net.LookupIP("www.baidu.com")
	fmt.Println(ips, err)

	ip, ipnet, err := net.ParseCIDR("192.168.1.1/24")
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.1.2")))
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.2.2")))
	fmt.Println(ipnet.Network())

	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		fmt.Println(addr.String(), addr.Network())
	}

	inters, _ := net.Interfaces()
	for _, inter := range inters {
		fmt.Println(inter.Addrs, inter.Index, inter.Name, inter.MTU, inter.HardwareAddr, inter.Flags)
		fmt.Println(inter.MulticastAddrs())
		fmt.Println(inter.Addrs())
	}
	fmt.Println(inters)

	// inters1, _ := net.InterfaceByIndex(0)
	// for _, inter := range inters1 {
	// 	fmt.Println(inter.Addrs, inter.Index, inter.Name, inter.MTU, inter.HardwareAddr, inter.Flags, inter.MulticastAddrs)
	// }
}
