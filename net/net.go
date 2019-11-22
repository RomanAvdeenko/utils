package net

import (
	"encoding/binary"
	"fmt"
	"github.com/RomanAvdeenko/utils/slice"
	"net"
)

func ExcludeInterfaces(in []net.Interface, excludeInterfaceNames []string) []net.Interface {
	// Get router interface names
	out := make([]net.Interface, 0, len(in))
	outNames := make([]string, 0, len(in))
	inNames := make([]string, 0, len(in))
	for _, i := range in {
		inNames = append(inNames, i.Name)
	}
	// Take needed interface names
	outNames = slice.DeleteElements(inNames, excludeInterfaceNames...)
	// Take needed interfaces
	for _, i := range in {
		for _, n := range outNames {
			if i.Name == n {
				out = append(out, i)
				break
			}
		}
	}
	return out
}

func Ip2int(ipAddr []byte) uint32 {
	fmt.Printf("%v, len:%v\n", ipAddr, len(ipAddr))
	if len(ipAddr) > net.IPv4len {
		return binary.BigEndian.Uint32(ipAddr[len(ipAddr)-net.IPv4len:])
	}
	return binary.BigEndian.Uint32(ipAddr)
}

func Broadcast(ipNet net.IPNet) net.IP {
	brd := make(net.IP, net.IPv4len)
	n := len(ipNet.IP)
	//fmt.Println(n)
	if n > net.IPv4len {
		for i := 0; i < net.IPv4len; i++ {
			brd[i] = ipNet.IP[n-net.IPv4len+i] | ^ipNet.Mask[i]
		}
		return brd
	}
	for i := 0; i < n; i++ {
		brd[i] = ipNet.IP[i] | ^ipNet.Mask[i]
	}
	return brd
}
