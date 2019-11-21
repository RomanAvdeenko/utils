package net

import (
	"encoding/binary"
	"errors"
	"github.com/RomanAvdeenko/utils/slice"
	"log"
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

func Ip2int(ipAddr string) (uint32, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, errors.New("wrong ipAddr format")
	}
	ip = ip.To4()
	log.Print(ip)
	return binary.BigEndian.Uint32(ip), nil
}
