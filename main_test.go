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
        {
            block{0x0DC53006, 0x00CAB840, 0xB38448E5, 0xE993F421},
            key{0x92BD9B1C, 0xE5D14101, 0x5445FBC9, 0x5E4D0EF2, 0x682080AA, 0x227D642F, 0x2687F934, 0x90405511},
            block{0xE12BDC1A, 0xE28257EC, 0x703FCCF0, 0x95EE8DF1},
        },
    }
    for _, test := range tests {
        if got := Fe(test.block, test.key); got != test.result {
            t.Errorf("Fe(%x, %x) is %x, not %x", test.block, test.key, test.result, got)
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
        {
            block{0x69CCA1C9, 0x3557C9E3, 0xD66BC3E0, 0xFA88FA6E},
            key{0xE9DEE72C, 0x8F0C0FA6, 0x2DDB49F4, 0x6F739647, 0x06075316, 0xED247A37, 0x39CBA383, 0x03A98BF6},
            block{0xB194BAC8, 0x0A08F53B, 0x366D008E, 0x584A5DE4},
        },
    }
    for _, test := range tests {
        if got := Fd(test.block, test.key); got != test.result {
            t.Errorf("Fd(%x, %x) is %x, not %x", test.block, test.key, test.result, got)
        }
    }
}

func equal(x, y []block) bool {
    if len(x) != len(y) {
        return false
    }
    for i := range x {
        if x[i] != y[i] {
            return false
        }
    }
    return true
}

func TestECBe(t *testing.T) {
    var tests = []struct{
        blocks []block
        key key
        result []block
    }{
        {
            []block{
                {0xB194BAC8, 0x0A08F53B, 0x366D008E, 0x584A5DE4},
                {0x8504FA9D, 0x1BB6C7AC, 0x252E72C2, 0x02FDCE0D},
                {0x5BE3D612, 0x17B96181, 0xFE6786AD, 0x716B890B},
            },
            key{0xE9DEE72C, 0x8F0C0FA6, 0x2DDB49F4, 0x6F739647, 0x06075316, 0xED247A37, 0x39CBA383, 0x03A98BF6},
            []block{
                {0x69CCA1C9, 0x3557C9E3, 0xD66BC3E0, 0xFA88FA6E},
                {0x5F23102E, 0xF1097107, 0x75017F73, 0x806DA9DC},
                {0x46FB2ED2, 0xCE771F26, 0xDCB5E5D1, 0x569F9AB0},
            },
        },
        {
            []block{
                {0x0DC53006, 0x00CAB840, 0xB38448E5, 0xE993F421},
                {0xE55A239F, 0x2AB5C5D5, 0xFDB6E81B, 0x40938E2A},
                {0x54120CA3, 0xE6E19C7A, 0xD750FC35, 0x31DAEAB7},
            },
            key{0x92BD9B1C, 0xE5D14101, 0x5445FBC9, 0x5E4D0EF2, 0x682080AA, 0x227D642F, 0x2687F934, 0x90405511},
            []block{
                {0xE12BDC1A, 0xE28257EC, 0x703FCCF0, 0x95EE8DF1},
                {0xC1AB7638, 0x9FE678CA, 0xF7C6F860, 0xD5BB9C4F},
                {0xF33C657B, 0x637C306A, 0xDD4EA779, 0x9EB23D31},
            },
        },
    }
    for _, test := range tests {
        if got := ECBe(test.blocks, test.key); !equal(got, test.result) {
            t.Errorf("ECBe(%x, %x) is %x, not %x", test.blocks, test.key, test.result, got)
        }
    }
}

func TestECBd(t *testing.T) {
    var tests = []struct{
        blocks []block
        key key
        result []block
    }{
        {
            []block{
                {0x69CCA1C9, 0x3557C9E3, 0xD66BC3E0, 0xFA88FA6E},
                {0x5F23102E, 0xF1097107, 0x75017F73, 0x806DA9DC},
                {0x46FB2ED2, 0xCE771F26, 0xDCB5E5D1, 0x569F9AB0},
            },
            key{0xE9DEE72C, 0x8F0C0FA6, 0x2DDB49F4, 0x6F739647, 0x06075316, 0xED247A37, 0x39CBA383, 0x03A98BF6},
            []block{
                {0xB194BAC8, 0x0A08F53B, 0x366D008E, 0x584A5DE4},
                {0x8504FA9D, 0x1BB6C7AC, 0x252E72C2, 0x02FDCE0D},
                {0x5BE3D612, 0x17B96181, 0xFE6786AD, 0x716B890B},
            },
        },
        {
            []block{
                {0xE12BDC1A, 0xE28257EC, 0x703FCCF0, 0x95EE8DF1},
                {0xC1AB7638, 0x9FE678CA, 0xF7C6F860, 0xD5BB9C4F},
                {0xF33C657B, 0x637C306A, 0xDD4EA779, 0x9EB23D31},
            },
            key{0x92BD9B1C, 0xE5D14101, 0x5445FBC9, 0x5E4D0EF2, 0x682080AA, 0x227D642F, 0x2687F934, 0x90405511},
            []block{
                {0x0DC53006, 0x00CAB840, 0xB38448E5, 0xE993F421},
                {0xE55A239F, 0x2AB5C5D5, 0xFDB6E81B, 0x40938E2A},
                {0x54120CA3, 0xE6E19C7A, 0xD750FC35, 0x31DAEAB7},
            },
        },
    }
    for _, test := range tests {
        if got := ECBd(test.blocks, test.key); !equal(got, test.result) {
            t.Errorf("ECBd(%x, %x) is %x, not %x", test.blocks, test.key, test.result, got)
        }
    }
}
