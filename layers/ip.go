package layers

import (
	"fmt"
	"net"
)

/* Usage */
//=================================================================================================//
// fmt.Println("全アドレス")
// GetLocalAllIP()
// fmt.Println()

// fmt.Println("全インターフェース")
// interfaces, err := GetLocalInterfaces()
// if err != nil {
// 		fmt.Println(err)
// }
// fmt.Println(interfaces)
// IPv4_list, IPv6_list := GetIPversion(interfaces)
// fmt.Println()

// fmt.Println("--IPv4--")
// GetLocalIP(IPv4_list)
// fmt.Println()

// fmt.Println("--IPv6--")
// GetLocalIP(IPv6_list)
// fmt.Println()

// fmt.Println("インターフェースに関連付けられたアドレス")
// Relevance(interfaces)
// fmt.Println()
//=================================================================================================//

// 全アドレス取得
func GetLocalAllIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok {
			fmt.Println(ipnet)
		}
	}
}

// 全インターフェース取得
func GetLocalInterfaces() ([]net.Interface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	return interfaces, nil
}

// バージョンごとに取得
func GetIPversion(interfaces []net.Interface) ([]string, []string) {
	var IPv4_list, IPv6_list []string

	for _, inter := range interfaces {
		addrs, err := inter.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok {
				if ipnet.IP.To4() != nil {
					// fmt.Println("IPv4:", ipnet.IP.String())
					IPv4_list = append(IPv4_list, ipnet.IP.String())
				} else {
					// fmt.Println("IPv6:", ipnet.IP.String())
					IPv6_list = append(IPv6_list, ipnet.IP.String())
				}
			}
		}
	}

	return IPv4_list, IPv6_list
}

// IPアドレス取得
func GetLocalIP(IP_list []string) {
	for _, ip := range IP_list {
		fmt.Println(ip)
	}
}

// インターフェースに紐付くIPアドレス
func Relevance(interfaces []net.Interface) {
	for _, interf := range interfaces {
		if addrs, err := interf.Addrs(); err == nil {
			for _, addr := range addrs {
				fmt.Println(interf.Name, ":", addr)
			}
		}
	}
}
