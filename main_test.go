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

func TestG(t *testing.T) {
    var tests = []struct{
        u u
        r uint32
        result u
    }{
        {
            u{0xb1, 0x94, 0xba, 0xc8},
            5,
            u{0x14, 0xa4, 0x3d, 0x1f},
        },
    }
    for _, test := range tests {
        if got := G(test.r, test.u); got != test.result {
            t.Errorf("G(%d, %x) is %x, not %x", test.r, test.u, test.result, got)
        }
    }
}
