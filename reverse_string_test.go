package reverse_string

import (
	"testing"
)

type testRevStr struct {
	name  string
	input string
	want  string
}

var testCases = []testRevStr{
	{"one", "Hello world!", "!dlrow olleH"},
	{"two", "Привет!", "!тевирП"},
}

func TestReverseString(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := ReverseString(testCase.input)

			if actual != testCase.want {
				t.Errorf("got: %s; want: %s", actual, testCase.want)
			}
		})
	}
}
