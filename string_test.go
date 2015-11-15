package input

import (
	"testing"
)

func TestTrimSingleWhitespaceNormalizeLine(t *testing.T) {
	for in, exp := range map[string]string{
		"Hello":                            "Hello",
		"   \r \n \t Hello \t":             "Hello",
		"   \r \n \t Hello      there \t":  "Hello there",
		"   \r \n \t Hello      there \t!": "Hello there !",
	} {
		act := TrimSingleWhitespaceNormalizeLine(in)
		if act != exp {
			t.Errorf("Expected TrimSingleWhitespaceNormalizeLine(`%s`) to return `%s` but it returned `%s`", in, exp, act)
		}
	}
}
