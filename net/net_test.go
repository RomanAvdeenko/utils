package net

import "testing"

var tests = []struct {
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
		[]byte{255, 255, 255, 0},
		4294967040,
	},
}

func TestIp2int(t *testing.T) {
	for _, test := range tests {
		if got := Ip2int(test.input); got != test.want {
			t.Error("input:", test.input, "got:", got, "want:", test.want)
		}
	}

}
