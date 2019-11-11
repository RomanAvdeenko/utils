package slice

import (
	"fmt"
	"reflect"
	"testing"
)

var tests = []struct {
	inputA []string
	inputB []string
	want   []string
}{
	{
		[]string{"lo", "bond0", "eth0", "eth1", "eth2", "eth3", "vlan113"},
		[]string{"lo", "bond0", "eth*"},
		[]string{"vlan113"},
	},
	{
		[]string{"lo", "eno1", "wlo1"},
		[]string{"lo*"},
		[]string{"eno1", "wlo1"},
	},
	{
		[]string{"Мама", "мыла", "раму", "и себя"},
		[]string{"раму", "и*"},
		[]string{"Мама", "мыла"},
	},
}

func TestDeleteElements(t *testing.T) {
	for _, test := range tests {
		if got := DeleteElements(test.inputA, test.inputB...); !reflect.DeepEqual(got, test.want) {
			t.Errorf("DeleteElements(%q, %q) returns %q want: %q ", test.inputA, test.inputB, got, test.want)
		}
	}

}

func BenchmarkDeleteElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DeleteElements(tests[0].inputA, tests[0].inputB...)
	}
}

func ExampleDeleteElements() {
	fmt.Println(DeleteElements([]string{"lo", "bond0", "eth0", "eth1", "eth2", "eth3", "vlan113"}, "lo", "bond0", "eth*"))
	fmt.Println(DeleteElements([]string{"lo", "eno1", "wlo1"}, "lo*"))
	// Output:
	// [vlan113]
	// [eno1 wlo1]
}
