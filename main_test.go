package main

import (
	"testing"
)

func TestH(t *testing.T) {
    var tests = []struct {
        input uint8
        output uint8
    }{
        {0xa2, 0x9b},
    }

    for _, test := range tests {
        if got := H(test.input); got != test.output {
            t.Errorf("H(%x) is %x, not %x", test.input, test.output, got)
        }
    }
}
