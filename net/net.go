package net

import (
	"encoding/binary"
	"errors"
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

// Возвращает uint32 предсьавление IP
func Ip2int(ipAddr []byte) uint32 {
	//fmt.Printf("%v, len:%v\n", ipAddr, len(ipAddr))
	if len(ipAddr) > net.IPv4len {
		return binary.BigEndian.Uint32(ipAddr[len(ipAddr)-net.IPv4len:])
	}
	return binary.BigEndian.Uint32(ipAddr)
}

// Возвращает []byte представление IP
func Ip2bytes(ip uint32) []byte {
	res := make([]byte, 4)
	res[0] = byte(ip)
	res[1] = byte(ip >> 8)
	res[2] = byte(ip >> 16)
	res[3] = byte(ip >> 24)
	return res
}

//Возващает brd IP адрес, CIDR-адреса
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

// Возващает IP адреса, CIDR-адреса за исключением  IP: сетиб brd и самого его
func IPs(addr net.Addr) ([]uint32, error) {
	var ips = []uint32{}
	ipNet, ok := addr.(*net.IPNet)
	if !ok {
		return []uint32{}, errors.New("getIP() type assertion error")
	}
	// Цикл по uint32 представлениям IP от netIP+1 до brd IP
	uint32Ip := Ip2int(ipNet.IP)
	uint32NetIp := Ip2int(ipNet.IP.Mask(ipNet.Mask))
	uint32Brd := Ip2int(Broadcast(*ipNet))
	//fmt.Printf("ip: %v\nnetIP: %v\nbrd ip: %v\n", uint32Ip, uint32NetIp, uint32Brd)
	n := uint32Brd - uint32NetIp + 1
	// /32 или /31 нет элементов
	if n < 4 {
		return []uint32{}, nil
	}
	// Уменишим размер на net ip brd
	n -= 3
	var c uint32
	ips = make([]uint32, n)
	for i := uint32NetIp + 1; i < uint32Brd; i++ {
		// Исключим ip интерфейса
		if i != uint32Ip {
			ips[c] = i
			c++
		}
	}
	return ips, nil
}
