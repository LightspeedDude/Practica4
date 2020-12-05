package code

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{001, 1},
		{1001, 11},
		{1011, 13},
	}
	for _, test := range tests {
		if output := cDtO(cBtD(test.input)); output != test.expected {
			t.Error("Test failed")
		}
	}
}
