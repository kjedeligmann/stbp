package main

var Hbox = [16][16]uint8{
    {0xB1, 0x94, 0xBA, 0xC8, 0x0A, 0x08, 0xF5, 0x3B, 0x36, 0x6D, 0x00, 0x8E, 0x58, 0x4A, 0x5D, 0xE4},
    {0x85, 0x04, 0xFA, 0x9D, 0x1B, 0xB6, 0xC7, 0xAC, 0x25, 0x2E, 0x72, 0xC2, 0x02, 0xFD, 0xCE, 0x0D},
    {0x5B, 0xE3, 0xD6, 0x12, 0x17, 0xB9, 0x61, 0x81, 0xFE, 0x67, 0x86, 0xAD, 0x71, 0x6B, 0x89, 0x0B},
    {0x5C, 0xB0, 0xC0, 0xFF, 0x33, 0xC3, 0x56, 0xB8, 0x35, 0xC4, 0x05, 0xAE, 0xD8, 0xE0, 0x7F, 0x99},
    {0xE1, 0x2B, 0xDC, 0x1A, 0xE2, 0x82, 0x57, 0xEC, 0x70, 0x3F, 0xCC, 0xF0, 0x95, 0xEE, 0x8D, 0xF1},
    {0xC1, 0xAB, 0x76, 0x38, 0x9F, 0xE6, 0x78, 0xCA, 0xF7, 0xC6, 0xF8, 0x60, 0xD5, 0xBB, 0x9C, 0x4F},
    {0xF3, 0x3C, 0x65, 0x7B, 0x63, 0x7C, 0x30, 0x6A, 0xDD, 0x4E, 0xA7, 0x79, 0x9E, 0xB2, 0x3D, 0x31},
    {0x3E, 0x98, 0xB5, 0x6E, 0x27, 0xD3, 0xBC, 0xCF, 0x59, 0x1E, 0x18, 0x1F, 0x4C, 0x5A, 0xB7, 0x93},
    {0xE9, 0xDE, 0xE7, 0x2C, 0x8F, 0x0C, 0x0F, 0xA6, 0x2D, 0xDB, 0x49, 0xF4, 0x6F, 0x73, 0x96, 0x47},
    {0x06, 0x07, 0x53, 0x16, 0xED, 0x24, 0x7A, 0x37, 0x39, 0xCB, 0xA3, 0x83, 0x03, 0xA9, 0x8B, 0xF6},
    {0x92, 0xBD, 0x9B, 0x1C, 0xE5, 0xD1, 0x41, 0x01, 0x54, 0x45, 0xFB, 0xC9, 0x5E, 0x4D, 0x0E, 0xF2},
    {0x68, 0x20, 0x80, 0xAA, 0x22, 0x7D, 0x64, 0x2F, 0x26, 0x87, 0xF9, 0x34, 0x90, 0x40, 0x55, 0x11},
    {0xBE, 0x32, 0x97, 0x13, 0x43, 0xFC, 0x9A, 0x48, 0xA0, 0x2A, 0x88, 0x5F, 0x19, 0x4B, 0x09, 0xA1},
    {0x7E, 0xCD, 0xA4, 0xD0, 0x15, 0x44, 0xAF, 0x8C, 0xA5, 0x84, 0x50, 0xBF, 0x66, 0xD2, 0xE8, 0x8A},
    {0xA2, 0xD7, 0x46, 0x52, 0x42, 0xA8, 0xDF, 0xB3, 0x69, 0x74, 0xC5, 0x51, 0xEB, 0x23, 0x29, 0x21},
    {0xD4, 0xEF, 0xD9, 0xB4, 0x3A, 0x62, 0x28, 0x75, 0x91, 0x14, 0x10, 0xEA, 0x77, 0x6C, 0xDA, 0x1D},
}

// The H substitution
func H(u uint8) uint8 {
    return Hbox[u >> 4][u & 0xF]
}

type block [4]uint32
type key [8]uint32

// The G transform
func G(r uint32, a uint32) (result uint32) {
    var u [4]uint8
    // Reading octets from uint32
    for i := 0; i < 4; i++ {
        u[i] = uint8(a >> ((3-i)*8))
    }
    // Using the H substitution on them
    for i := 0; i < 4; i++ {
        u[i] = H(u[i])
    }
    // Putting the result into uint32 in little-endian order
    var leu uint32
    for i := 0; i < 4; i++ {
        leu |= uint32(u[i]) << (i*8)
    }
    // Cyclic bit shift
    leu = leu << r | leu >> (32-r)
    // Putting the octets back in the big-endian order
    for i := 0; i < 4; i++ {
        result |= uint32(uint8(leu)) << ((3-i)*8)
        leu >>= 8
    }
    return
}

// Little-endian plus from the standard
func Plus(u uint32, v uint32) (r uint32) {
    var c uint32
    for i := 0; i < 4; i++ {
        c |= uint32(uint8(u)) << ((3-i)*8)
        u >>= 8
    }
    r = c
    c = 0
    for i := 0; i < 4; i++ {
        c |= uint32(uint8(v)) << ((3-i)*8)
        v >>= 8
    }
    r += c
    c = 0
    for i := 0; i < 4; i++ {
        c |= uint32(uint8(r)) << ((3-i)*8)
        r >>= 8
    }
    r = c
    return
}

// Little-endian minus from the standard
func Minus(u uint32, v uint32) (r uint32) {
    var c uint32
    for i := 0; i < 4; i++ {
        c |= uint32(uint8(u)) << ((3-i)*8)
        u >>= 8
    }
    r = c
    c = 0
    for i := 0; i < 4; i++ {
        c |= uint32(uint8(v)) << ((3-i)*8)
        v >>= 8
    }
    r -= c
    c = 0
    for i := 0; i < 4; i++ {
        c |= uint32(uint8(r)) << ((3-i)*8)
        r >>= 8
    }
    r = c
    return
}

// Encryption function (F_theta(x))
func Fe(x block, k key) block {
    var a, b, c, d, e uint32
    a, b, c, d = x[0], x[1], x[2], x[3]
    for i := 1; i <= 8; i++ {
        b ^= G(5, Plus(a, k[((7*i-6)-1)%8]))
        c ^= G(21, Plus(d, k[((7*i-5)-1)%8]))
        a = Minus(a, G(13, Plus(b, k[((7*i-4)-1)%8])))
        // i << 24 because of little-endianness in the standard
        e = G(21, Plus(Plus(b, c), k[((7*i-3)-1)%8])) ^ uint32(i << 24)
        b = Plus(b, e)
        c = Minus(c, e)
        d = Plus(d, G(13, Plus(c, k[((7*i-2)-1)%8])))
        b ^= G(21, Plus(a, k[((7*i-1)-1)%8]))
        c ^= G(5, Plus(d, k[((7*i)-1)%8]))
        a, b = b, a
        c, d = d, c
        b, c = c, b
    }
    y := block{b, d, a, c}
    return y
}

// Decryption function (F_theta^{-1}(x))
func Fd(x block, k key) block {
    var a, b, c, d, e uint32
    a, b, c, d = x[0], x[1], x[2], x[3]
    for i := 8; i >= 1; i-- {
        b ^= G(5, Plus(a, k[((7*i)-1)%8]))
        c ^= G(21, Plus(d, k[((7*i-1)-1)%8]))
        a = Minus(a, G(13, Plus(b, k[((7*i-2)-1)%8])))
        // i << 24 because of little-endianness in the standard
        e = G(21, Plus(Plus(b, c), k[((7*i-3)-1)%8])) ^ uint32(i << 24)
        b = Plus(b, e)
        c = Minus(c, e)
        d = Plus(d, G(13, Plus(c, k[((7*i-4)-1)%8])))
        b ^= G(21, Plus(a, k[((7*i-5)-1)%8]))
        c ^= G(5, Plus(d, k[((7*i-6)-1)%8]))
        a, b = b, a
        c, d = d, c
        a, d = d, a
    }
    y := block{c, a, d, b}
    return y
}

// ECB and CBC modes will work with slices of blocks for now (but this is not optimal)

// ECB mode
func ECBe(x []block, k key) (y []block) {
    for i := 0; i < len(x); i++ {
        y = append(y, Fe(x[i], k))
    }
    return
}

func ECBd(x []block, k key) (y []block) {
    for i := 0; i < len(x); i++ {
        y = append(y, Fd(x[i], k))
    }
    return
}

// CBC mode
func xb(x, y block) (z block) {
    for i := 0; i < 4; i++ {
        z[i] = x[i] ^ y[i]
    }
    return
}

func CBCe(x []block, k key, s block) (y []block) {
    y0 := Fe(s, k)
    y = append(y, Fe(xb(x[0], y0), k))
    for i := 1; i < len(x); i++ {
        y = append(y, Fe(xb(x[i], y[i-1]), k))
    }
    return
}

func CBCd(x []block, k key, s block) (y []block) {
    x0 := Fe(s, k)
    y = append(y, xb(Fd(x[0], k), x0))
    for i := 1; i < len(x); i++ {
        y = append(y, xb(Fd(x[i], k), x[i-1]))
    }
    return
}

// CFB mode
func CFBe(x []byte, k key, s block) (y []byte) {
    for i := 0; i < len(x); i += 16 {
        s = Fe(s, k)
        for j := 0; j < 16 && i+j < len(x); j++ {
            s[j/4] ^= (uint32(x[i+j]) << ((3-(j%4))*8))
            y = append(y, byte(s[j/4] >> ((3-(j%4))*8)))
        }
    }
    return
}

func CFBd(x []byte, k key, s block) (y []byte) {
    var s1 block
    for i := 0; i < len(x); i += 16 {
        s = Fe(s, k)
        s1 = s
        for j := 0; j < 16 && i+j < len(x); j++ {
            s[j/4] ^= (uint32(x[i+j]) << ((3-(j%4))*8))
            y = append(y, byte(s[j/4] >> ((3-(j%4))*8)))
        }
        s = xb(s, s1)
    }
    return
}

// CTR (Counter) mode
func (s *block) Increment() {
    for i := 0; i < 4; i++ {
        s[i] = Plus(s[i], uint32(1 << 24))
        if s[i] != 0 {
            break
        }
    }
}

func CTR(x []byte, k key, s block) (y []byte) {
    s = Fe(s, k)
    for i := 0; i < len(x); i += 16 {
        s.Increment()
        es := Fe(s, k)
        for j := 0; j < 16 && i+j < len(x); j++ {
            y = append(y, x[i+j] ^ byte(es[j/4] >> ((3-(j%4))*8)))
        }
    }
    return
}
