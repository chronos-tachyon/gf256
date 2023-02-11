package gf256

// GF256 represents a GF(256) field.
type GF256 struct {
	exp [512]byte
	log [256]byte
	p   Polynomial
}

// New constructs a GF256 value representing a Galois Field with multiplication
// defined in relation to the given irreducible polynomial with binary
// coefficients.  The returned GF256 object is immutable and fully thread-safe.
func New(p Polynomial) *GF256 {
	p.checkIndex()

	field := new(GF256)
	field.p = p
	x := byte(1)
	for i := uint(0); i < 255; i++ {
		field.exp[i] = x
		field.log[x] = byte(i)
		x = slowMul(x, 2, p)
	}
	for i := uint(255); i < 512; i++ {
		field.exp[i] = field.exp[i-255]
	}
	return field
}

// Polynomial returns the identifying enum constant for the irreducible
// polynomial chosen to define this field.
func (field *GF256) Polynomial() Polynomial {
	return field.p
}

// Add returns the sum of two GF(256) elements.
func (field *GF256) Add(a byte, b byte) byte {
	return a ^ b
}

// Sub returns the difference of two GF(256) elements.  In all GF(2**k) fields,
// addition and subtraction are the same operation because each element is its
// own additive inverse.
func (field *GF256) Sub(a byte, b byte) byte {
	return field.Add(a, b)
}

// Mul returns the product of two GF(256) elements, modulo the irreducible
// polynomial which was used to construct the field.
func (field *GF256) Mul(a byte, b byte) byte {
	if a == 0 || b == 0 {
		return 0
	}
	logA := uint(field.log[a])
	logB := uint(field.log[b])
	return field.exp[logA+logB]
}

// Div returns the divisor of two GF(256) elements, modulo the irreducible
// polynomial which was used to construct the field.  As usual, division is
// only defined when the divisor is non-zero.
func (field *GF256) Div(a byte, b byte) (byte, bool) {
	if b == 0 {
		return 0, false
	}
	if a == 0 {
		return 0, true
	}
	logA := uint(field.log[a])
	logB := uint(field.log[b])
	return field.exp[(logA+255-logB)%255], true
}

// Inv returns the multiplicative inverse of a non-zero GF(256) element.  In
// other words, Mul(a, Inv(b)) is the same thing as Div(a, b).
func (field *GF256) Inv(a byte) (byte, bool) {
	if a == 0 {
		return 0, false
	}
	logA := uint(field.log[a])
	return field.exp[255-logA], true
}

// Pow returns the first value, which must be an element of GF(256), to the
// power of the second value.
func (field *GF256) Pow(a byte, b uint) byte {
	if a == 0 {
		return 0
	}
	logA := uint(field.log[a])
	return field.exp[(logA*b)%255]
}
