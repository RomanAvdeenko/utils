package net

import (
	"bytes"
	"net"
	"testing"
)

var testIp2int = []struct {
	input []byte
	want  uint32
}{
	{
		[]byte{193, 111, 156, 1},
		3245317121,
	},
	{
		[]byte{192, 168, 1, 203},
		3232235979,
	},
	{
		[]byte{127, 0, 0, 1},
		2130706433,
	},
	{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 0},
		4294967040,
	},
}

var testBroadcast = []struct {
	input net.IPNet
	want  net.IP
}{
	{
		net.IPNet{
			[]byte{193, 111, 156, 10}, []byte{255, 255, 255, 224},
		},
		[]byte{193, 111, 156, 31},
	},
	{
		net.IPNet{
			[]byte{91, 192, 159, 2}, []byte{255, 255, 255, 240},
		},
		[]byte{91, 192, 159, 15},
	},
	{
		net.IPNet{
			[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 192, 168, 0, 24}, []byte{255, 255, 255, 0},
		},
		[]byte{192, 168, 0, 255},
	},
}

var testIp2bytes = []struct {
	input uint32
	want  []byte
}{
	{
		3245317122,
		[]byte{193, 111, 156, 2},
	},
}

func TestIp2int(t *testing.T) {
	for _, test := range testIp2int {
		if got := Ip2int(test.input); got != test.want {
			t.Error("input:", test.input, "got:", got, "want:", test.want)
		}
	}
}
func TestIp2bytes(t *testing.T) {
	for _, test := range testIp2bytes {
		if got := Ip2bytes(test.input); !bytes.Equal(got, test.want) {
			t.Error("input:", test.input, "got:", got, "want:", test.want)
		}
	}
}
func TestBroadcast(t *testing.T) {
	for _, test := range testBroadcast {
		if got := Broadcast(test.input); !bytes.Equal(got, test.want) {
			t.Error("input:", test.input, "got:", got, "want:", test.want)
		}
	}

}
