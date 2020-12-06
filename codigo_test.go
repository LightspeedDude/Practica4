package codigo

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{001, 1},
		{1001, 11},
		{101010, 52},
	}
	for _, test := range tests {
		if output := cDaO(cBaD(test.input)); output != test.expected {
			t.Error("test failed")
		}
	}
}
