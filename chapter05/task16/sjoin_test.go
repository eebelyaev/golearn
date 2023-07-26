package sjoin

import (
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		inputSep string
		input    []string
	}{
		{";", []string{"a", "b", "c"}},
		{";", []string{}},
		{" + ", []string{"a", "b", "c"}},
	}
	for _, test := range tests {
		res := Join(test.inputSep, test.input...)
		wait := strings.Join(test.input, test.inputSep)
		if res != wait {
			t.Errorf("wait: %s, got: %s", wait, res)
		}
	}
}
