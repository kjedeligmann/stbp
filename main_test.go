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
        u uint32
        r uint32
        result uint32
    }{
        {
            0xb194bac8,
            5,
            0x14a43d1f,
        },
    }
    for _, test := range tests {
        if got := G(test.r, test.u); got != test.result {
            t.Errorf("G(%d, %x) is %x, not %x", test.r, test.u, test.result, got)
        }
    }
}

func TestFe(t *testing.T) {
    var tests = []struct{
        block block
        key key
        result block
    }{
        {
            block{0xB194BAC8, 0x0A08F53B, 0x366D008E, 0x584A5DE4},
            key{0xE9DEE72C, 0x8F0C0FA6, 0x2DDB49F4, 0x6F739647, 0x06075316, 0xED247A37, 0x39CBA383, 0x03A98BF6},
            block{0x69CCA1C9, 0x3557C9E3, 0xD66BC3E0, 0xFA88FA6E},
        },
    }
    for _, test := range tests {
        if got := Fe(test.block, test.key); got != test.result {
            t.Errorf("ECBe(%x, %x) is %x, not %x", test.block, test.key, test.result, got)
        }
    }
}

func TestFd(t *testing.T) {
    var tests = []struct{
        block block
        key key
        result block
    }{
        {
            block{0xE12BDC1A, 0xE28257EC, 0x703FCCF0, 0x95EE8DF1},
            key{0x92BD9B1C, 0xE5D14101, 0x5445FBC9, 0x5E4D0EF2, 0x682080AA, 0x227D642F, 0x2687F934, 0x90405511},
            block{0x0DC53006, 0x00CAB840, 0xB38448E5, 0xE993F421},
        },
    }
    for _, test := range tests {
        if got := Fd(test.block, test.key); got != test.result {
            t.Errorf("ECBd(%x, %x) is %x, not %x", test.block, test.key, test.result, got)
        }
    }
}
