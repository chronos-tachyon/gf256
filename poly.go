package gf256

import "fmt"

// Polynomial is an enumeration constant that selects one of the 16 irreducible
// polynomials that exist in GF(256).
type Polynomial byte

const (
	// Poly11D represents the 8th degree polynomial with coefficients 1 0001 1101.
	Poly11D Polynomial = iota
	// Poly12B represents the 8th degree polynomial with coefficients 1 0010 1011.
	Poly12B
	// Poly12D represents the 8th degree polynomial with coefficients 1 0010 1101.
	Poly12D
	// Poly14D represents the 8th degree polynomial with coefficients 1 0100 1101.
	Poly14D
	// Poly15F represents the 8th degree polynomial with coefficients 1 0101 1111.
	Poly15F
	// Poly163 represents the 8th degree polynomial with coefficients 1 0110 0011.
	Poly163
	// Poly165 represents the 8th degree polynomial with coefficients 1 0110 0101.
	Poly165
	// Poly169 represents the 8th degree polynomial with coefficients 1 0110 1001.
	Poly169
	// Poly171 represents the 8th degree polynomial with coefficients 1 0111 0001.
	Poly171
	// Poly187 represents the 8th degree polynomial with coefficients 1 1000 0111.
	Poly187
	// Poly18C represents the 8th degree polynomial with coefficients 1 1000 1100.
	Poly18C
	// Poly1A9 represents the 8th degree polynomial with coefficients 1 1010 1001.
	Poly1A9
	// Poly1C3 represents the 8th degree polynomial with coefficients 1 1100 0011.
	Poly1C3
	// Poly1CF represents the 8th degree polynomial with coefficients 1 1100 1111.
	Poly1CF
	// Poly1E7 represents the 8th degree polynomial with coefficients 1 1110 0111.
	Poly1E7
	// Poly1F5 represents the 8th degree polynomial with coefficients 1 1111 0101.
	Poly1F5
)

// NumPolynomials is the number of irreducible polynomials.  All Polynomial
// enum constants are in the range from 0 (inclusive) to this value
// (exclusive).
const NumPolynomials = 16

var polyGoNames = [16]string{
	"gf256.Poly11D",
	"gf256.Poly12B",
	"gf256.Poly12D",
	"gf256.Poly14D",
	"gf256.Poly15F",
	"gf256.Poly163",
	"gf256.Poly165",
	"gf256.Poly169",
	"gf256.Poly171",
	"gf256.Poly187",
	"gf256.Poly18C",
	"gf256.Poly1A9",
	"gf256.Poly1C3",
	"gf256.Poly1CF",
	"gf256.Poly1E7",
	"gf256.Poly1F5",
}

var polyNames = [16]string{
	"0x11d",
	"0x12b",
	"0x12d",
	"0x14d",
	"0x15f",
	"0x163",
	"0x165",
	"0x169",
	"0x171",
	"0x187",
	"0x18c",
	"0x1a9",
	"0x1c3",
	"0x1cf",
	"0x1e7",
	"0x1f5",
}

var polyArray = [16]byte{
	0x1d, // 1 0001 1101  0x11d
	0x2b, // 1 0010 1011  0x12b
	0x2d, // 1 0010 1101  0x12d
	0x4d, // 1 0100 1101  0x14d
	0x5f, // 1 0101 1111  0x15f
	0x63, // 1 0110 0011  0x163
	0x65, // 1 0110 0101  0x165
	0x69, // 1 0110 1001  0x169
	0x71, // 1 0111 0001  0x171
	0x87, // 1 1000 0111  0x187
	0x8c, // 1 1000 1101  0x18c
	0xa9, // 1 1010 1001  0x1a9
	0xc3, // 1 1100 0011  0x1c3
	0xcf, // 1 1100 1111  0x1cf
	0xe7, // 1 1110 0111  0x1e7
	0xf5, // 1 1111 0101  0x1f5
}

func (p Polynomial) checkIndex() {
	if p >= NumPolynomials {
		panic(fmt.Errorf("Polynomial index %d is out of range [0 .. 15]", p))
	}
}

// GoString returns the name of the polynomial's Go constant.
func (p Polynomial) GoString() string {
	p.checkIndex()
	return polyGoNames[p]
}

// String returns the name of the polynomial, which is simply Value() expressed
// in hexadecimal format.
func (p Polynomial) String() string {
	p.checkIndex()
	return polyNames[p]
}

// Value returns the polynomial's coefficients as a 9-bit integer.
func (p Polynomial) Value() uint {
	p.checkIndex()
	return uint(polyArray[p]) | 0x100
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
