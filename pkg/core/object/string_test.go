package object

import (
	"testing"
)

func TestString_Inspect(t *testing.T) {
	testcases := []struct {
		obj      *String
		expected string
	}{
		{NewString(`Abcde`), `"Abcde"`},
		{NewString(`あいうえお`), `"あいうえお"`},
		{NewString(`タブ	入り`), `"タブ\t入り"`},
		{NewString(`"ダブルクォート" 入り`), `"\"ダブルクォート\" 入り"`},
	}

	for _, test := range testcases {
		t.Run(test.expected, func(t *testing.T) {
			actual := test.obj.Inspect()
			if actual != test.expected {
				t.Fatalf("got=%q, want=%q", actual, test.expected)
			}
		})
	}
}
