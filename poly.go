package gf256

import (
	"fmt"
	"strconv"
)

// NumPolynomials is the number of irreducible polynomials.  All Polynomial
// enum constants are in the range from 0 (inclusive) to this value
// (exclusive).
const NumPolynomials = 30

// Polynomial is an enumeration constant that selects one of the 30 monic
// irreducible polynomials that are suitable for constructing GF(256).
//
// Ref: https://codyplanteen.com/assets/rs/gf256_prim.pdf
type Polynomial byte

const (
	Poly11B Polynomial = iota
	Poly11D
	Poly12B
	Poly12D
	Poly139
	Poly13F
	Poly14D
	Poly15F
	Poly163
	Poly165
	Poly169
	Poly171
	Poly177
	Poly17B
	Poly187
	Poly18B
	Poly18D
	Poly19F
	Poly1A3
	Poly1A9
	Poly1B1
	Poly1BD
	Poly1C3
	Poly1CF
	Poly1D7
	Poly1DD
	Poly1E7
	Poly1F3
	Poly1F5
	Poly1F9
)

var polyGoNames = [NumPolynomials]string{
	"gf256.Poly11B",
	"gf256.Poly11D",
	"gf256.Poly12B",
	"gf256.Poly12D",
	"gf256.Poly139",
	"gf256.Poly13F",
	"gf256.Poly14D",
	"gf256.Poly15F",
	"gf256.Poly163",
	"gf256.Poly165",
	"gf256.Poly169",
	"gf256.Poly171",
	"gf256.Poly177",
	"gf256.Poly17B",
	"gf256.Poly187",
	"gf256.Poly18B",
	"gf256.Poly18D",
	"gf256.Poly19F",
	"gf256.Poly1A3",
	"gf256.Poly1A9",
	"gf256.Poly1B1",
	"gf256.Poly1BD",
	"gf256.Poly1C3",
	"gf256.Poly1CF",
	"gf256.Poly1D7",
	"gf256.Poly1DD",
	"gf256.Poly1E7",
	"gf256.Poly1F3",
	"gf256.Poly1F5",
	"gf256.Poly1F9",
}

var polyArray = [NumPolynomials]byte{
	0x1b, // 11b = 1 0001 1011 = x^8 + x^4 + x^3 + x + 1
	0x1d, // 11d = 1 0001 1101 = x^8 + x^4 + x^3 + x^2 + 1
	0x2b, // 12b = 1 0010 1011 = x^8 + x^5 + x^3 + x + 1
	0x2d, // 12d = 1 0010 1101 = x^8 + x^5 + x^3 + x^2 + 1
	0x39, // 139 = 1 0011 1001 = x^8 + x^5 + x^4 + x^3 + 1
	0x3f, // 13f = 1 0011 1111 = x^8 + x^5 + x^4 + x^3 + x^2 + x + 1
	0x4d, // 14d = 1 0100 1101 = x^8 + x^6 + x^3 + x^2 + 1
	0x5f, // 15f = 1 0101 1111 = x^8 + x^6 + x^4 + x^3 + x^2 + x + 1
	0x63, // 163 = 1 0110 0011 = x^8 + x^6 + x^5 + x + 1
	0x65, // 165 = 1 0110 0101 = x^8 + x^6 + x^5 + x^2 + 1
	0x69, // 169 = 1 0110 1001 = x^8 + x^6 + x^5 + x^3 + 1
	0x71, // 171 = 1 0111 0001 = x^8 + x^6 + x^5 + x^4 + 1
	0x77, // 177 = 1 0111 0111 = x^8 + x^6 + x^5 + x^4 + x^2 + x + 1
	0x7b, // 17b = 1 0111 1011 = x^8 + x^6 + x^5 + x^4 + x^3 + x + 1
	0x87, // 187 = 1 1000 0111 = x^8 + x^7 + x^2 + x + 1
	0x8b, // 18b = 1 1000 1011 = x^8 + x^7 + x^3 + x + 1
	0x8d, // 18d = 1 1000 1101 = x^8 + x^7 + x^3 + x^2 + 1
	0x9f, // 19f = 1 1001 1111 = x^8 + x^7 + x^4 + x^3 + x^2 + x + 1
	0xa3, // 1a3 = 1 1010 0011 = x^8 + x^7 + x^5 + x + 1
	0xa9, // 1a9 = 1 1010 1001 = x^8 + x^7 + x^5 + x^3 + 1
	0xb1, // 1b1 = 1 1011 0001 = x^8 + x^7 + x^5 + x^4 + 1
	0xbd, // 1bd = 1 1011 1101 = x^8 + x^7 + x^5 + x^4 + x^3 + x^2 + 1
	0xc3, // 1c3 = 1 1100 0011 = x^8 + x^7 + x^6 + x + 1
	0xcf, // 1cf = 1 1100 1111 = x^8 + x^7 + x^6 + x^3 + x^2 + x + 1
	0xd7, // 1d7 = 1 1101 0111 = x^8 + x^7 + x^6 + x^4 + x^2 + x + 1
	0xdd, // 1dd = 1 1101 1101 = x^8 + x^7 + x^6 + x^4 + x^3 + x^2 + 1
	0xe7, // 1e7 = 1 1110 0111 = x^8 + x^7 + x^6 + x^5 + x^2 + x + 1
	0xf3, // 1f3 = 1 1111 0011 = x^8 + x^7 + x^6 + x^5 + x^4 + x + 1
	0xf5, // 1f5 = 1 1111 0101 = x^8 + x^7 + x^6 + x^5 + x^4 + x^2 + 1
	0xf9, // 1f9 = 1 1111 1001 = x^8 + x^7 + x^6 + x^5 + x^4 + x^3 + 1
}

var polyGenerator = [NumPolynomials]byte{
	3, 2, 2, 2, 3,
	3, 2, 2, 2, 2,
	2, 2, 3, 9, 2,
	6, 2, 3, 3, 2,
	6, 7, 2, 2, 7,
	6, 2, 6, 2, 3,
}

func (p Polynomial) checkIndex() {
	if p >= NumPolynomials {
		panic(fmt.Errorf("Polynomial index %d is out of range [0 .. %d]", p, NumPolynomials))
	}
}

// GoString returns the name of the polynomial's Go constant.
func (p Polynomial) GoString() string {
	p.checkIndex()
	return polyGoNames[p]
}

// String returns the polynomial as a string.
func (p Polynomial) String() string {
	var tmp [64]byte
	slice := tmp[:0]

	needSep := false
	term := func(degree uint) {
		if needSep {
			slice = append(slice, '+')
		}
		switch {
		case degree > 1:
			slice = append(slice, 'x', '^')
			slice = strconv.AppendUint(slice, uint64(degree), 10)
		case degree == 1:
			slice = append(slice, 'x')
		case degree == 0:
			slice = append(slice, '1')
		}
		needSep = true
	}

	value := p.Value()
	degree := uint(9)
	for degree > 0 {
		degree--
		bit := uint(1) << degree
		if (value & bit) != 0 {
			term(degree)
		}
	}
	return string(slice)
}

// Value returns the polynomial's coefficients as a 9-bit integer.
func (p Polynomial) Value() uint {
	p.checkIndex()
	return uint(polyArray[p]) | 0x100
}

// Generator returns the generator element for this polynomial.
func (p Polynomial) Generator() byte {
	p.checkIndex()
	return polyGenerator[p]
}

func slowMul(a byte, b byte, p Polynomial) byte {
	polyBits := polyArray[p]
	var c byte
	for b != 0 {
		lowBitOfB := (b & 1) != 0
		highBitOfA := (a & 0x80) != 0
		if lowBitOfB {
			c ^= a
		}
		b >>= 1
		a <<= 1
		if highBitOfA {
			a ^= polyBits
		}
	}
	return c
}
