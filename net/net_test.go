package net

import "testing"

var tests = []struct {
	input string
	want  uint32
}{
	{
		"193.111.156.1",
		3245317121,
	},
	{
		"192.168.1.203",
		3232235979,
	},
	{
		"127.0.0.1",
		2130706433,
	},
}

func TestIp2int(t *testing.T) {
	for _, test := range tests {
		if got, _ := Ip2int(test.input); got != test.want {
			t.Error("input:", test.input, "got:", got, "want:", test.want)
		}
	}

}
