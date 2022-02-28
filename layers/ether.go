package layers

import (
	"fmt"
	"net"
)

const (
	EthTypeArp  uint16 = 0x0806
	EthTypeIpv4 uint16 = 0x0800
	EthTypeIpv6 uint16 = 0x86dd
)

type EtherHeader struct {
	DstMacAddr net.HardwareAddr
	SrcMacAddr net.HardwareAddr
	ProtoType  uint16
}

/* Usage */
//=================================================================================================//
// var interf string
// var netInterface *net.Interface
// var err error

// // interf = "en0"
// interf = "eth0"
// netInterface, err = CheckInterface(interf)
// if err != nil {
// 	fmt.Println(err)
// }
// GetHardwareAddr(netInterface)

// // interf = "en10"
// interf = "eth1"
// netInterface, err = CheckInterface(interf)
// if err != nil {
// 	fmt.Println(err)
// }

// fmt.Println(netInterface.Name)
// fmt.Println(netInterface.HardwareAddr)

// GetHardwareAddr(netInterface)
//=================================================================================================//

// CheckInterface returns the obtained interface information
func CheckInterface(interf string) (netInterface *net.Interface, err error) {
	for {
		netInterface, err = net.InterfaceByName(interf)
		if err != nil {
			fmt.Println("Can't find interface")
		} else {
			break
		}
	}
	fmt.Println("Connected.")

	return netInterface, nil
}

// GetHardwareAddr retruns extract the hardware information base on the interface name
func GetHardwareAddr(netInterface *net.Interface) {
	name := netInterface.Name
	macAddress := netInterface.HardwareAddr

	fmt.Println("Hardware name :", name)
	fmt.Println("MAC address :", macAddress)

	// verify if the MAC address can be parsed properly
	hwAddr, err := net.ParseMAC(macAddress.String())
	if err != nil {
		fmt.Println("No able to parse MAC address :", err)
	}
	fmt.Printf("Physical hardware address :%s \n", hwAddr.String())
}
