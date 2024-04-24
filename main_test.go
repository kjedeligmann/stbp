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

func TestCBCe(t *testing.T) {
    var tests = []struct{
        blocks []block
        key key
        iv block
        result []block
    }{
        {
            []block{
                {0xB194BAC8, 0x0A08F53B, 0x366D008E, 0x584A5DE4},
                {0x8504FA9D, 0x1BB6C7AC, 0x252E72C2, 0x02FDCE0D},
                {0x5BE3D612, 0x17B96181, 0xFE6786AD, 0x716B890B},
            },
            key{0xE9DEE72C, 0x8F0C0FA6, 0x2DDB49F4, 0x6F739647, 0x06075316, 0xED247A37, 0x39CBA383, 0x03A98BF6},
            block{0xBE329713, 0x43FC9A48, 0xA02A885F, 0x194B09A1},
            []block{
                {0x10116EFA, 0xE6AD58EE, 0x14852E11, 0xDA1B8A74},
                {0x5CF2480E, 0x8D03F1C1, 0x9492E53E, 0xD3A70F60},
                {0x657C1EE8, 0xC0E0AE5B, 0x58388BF8, 0xA68E3309},
            },
        },
        {
            []block{
                {0x730894D6, 0x158E17CC, 0x1600185A, 0x8F411CAB},
                {0x0471FF85, 0xC8379239, 0x8D8924EB, 0xD57D03DB},
                {0x95B97A9B, 0x7907E4B0, 0x20960455, 0xE46176F8},
            },
            key{0x92BD9B1C, 0xE5D14101, 0x5445FBC9, 0x5E4D0EF2, 0x682080AA, 0x227D642F, 0x2687F934, 0x90405511},
            block{0x7ECDA4D0, 0x1544AF8C, 0xA58450BF, 0x66D2E88A},
            []block{
                {0xE12BDC1A, 0xE28257EC, 0x703FCCF0, 0x95EE8DF1},
                {0xC1AB7638, 0x9FE678CA, 0xF7C6F860, 0xD5BB9C4F},
                {0xF33C657B, 0x637C306A, 0xDD4EA779, 0x9EB23D31},
            },
        },
    }
    for _, test := range tests {
        // In 2007 edition of the standard Y0, X0 are equal to F_theta(S), but in 2011 it is changed to plain S. Testcases are from the 2011 edition
        if got := CBCe(test.blocks, test.key, Fd(test.iv, test.key)); !equal(got, test.result) {
            t.Errorf("CBCe(%x, %x, %x) is %x, not %x", test.blocks, test.key, test.iv, test.result, got)
        }
    }
}

func TestCBCd(t *testing.T) {
    var tests = []struct{
        blocks []block
        key key
        iv block
        result []block
    }{
        {
            []block{
                {0x10116EFA, 0xE6AD58EE, 0x14852E11, 0xDA1B8A74},
                {0x5CF2480E, 0x8D03F1C1, 0x9492E53E, 0xD3A70F60},
                {0x657C1EE8, 0xC0E0AE5B, 0x58388BF8, 0xA68E3309},
            },
            key{0xE9DEE72C, 0x8F0C0FA6, 0x2DDB49F4, 0x6F739647, 0x06075316, 0xED247A37, 0x39CBA383, 0x03A98BF6},
            block{0xBE329713, 0x43FC9A48, 0xA02A885F, 0x194B09A1},
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
            block{0x7ECDA4D0, 0x1544AF8C, 0xA58450BF, 0x66D2E88A},
            []block{
                {0x730894D6, 0x158E17CC, 0x1600185A, 0x8F411CAB},
                {0x0471FF85, 0xC8379239, 0x8D8924EB, 0xD57D03DB},
                {0x95B97A9B, 0x7907E4B0, 0x20960455, 0xE46176F8},
            },
        },
    }
    for _, test := range tests {
        // Fd() is used because of the later change in the standard
        if got := CBCd(test.blocks, test.key, Fd(test.iv, test.key)); !equal(got, test.result) {
            t.Errorf("CBCd(%x, %x, %x) is %x, not %x", test.blocks, test.key, test.iv, test.result, got)
        }
    }
}
