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
func (field *GF256) Add(x byte, y byte) byte {
	return x ^ y
}

// Sub returns the difference of two GF(256) elements.  In all GF(2**k) fields,
// addition and subtraction are the same operation because each element is its
// own additive inverse.
func (field *GF256) Sub(x byte, y byte) byte {
	return field.Add(x, y)
}

// Mul returns the product of two GF(256) elements, modulo the irreducible
// polynomial which was used to construct the field.
func (field *GF256) Mul(x byte, y byte) byte {
	if x == 0 || y == 0 {
		return 0
	}
	logX := uint(field.log[x])
	logY := uint(field.log[y])
	return field.exp[logX+logY]
}

// Div returns the divisor of two GF(256) elements, modulo the irreducible
// polynomial which was used to construct the field.  As usual, division is
// only defined when the divisor is non-zero.
func (field *GF256) Div(x byte, y byte) (byte, bool) {
	if y == 0 {
		return 0, false
	}
	if x == 0 {
		return 0, true
	}
	logX := uint(field.log[x]) + 255
	logY := uint(field.log[y])
	return field.exp[(logX-logY)%255], true
}

// Inv returns the multiplicative inverse of a non-zero GF(256) element.  In
// other words, Mul(x, Inv(y)) is the same thing as Div(x, y).
func (field *GF256) Inv(x byte) (byte, bool) {
	if x == 0 {
		return 0, false
	}
	return field.exp[255-field.log[x]], true
}

// Pow returns the first value, which must be an element of GF(256), to the
// power of the second value.
func (field *GF256) Pow(x byte, y uint) byte {
	if x == 0 {
		return 0
	}
	logX := uint(field.log[x])
	return field.exp[(logX*y)%255]
}

// Log returns the base-2 logarithm of its argument, with multiplication
// defined modulo the irreducible polynomial with which this field was
// constructed.  As usual, the logarithm of 0 is undefined.
func (field *GF256) Log(x byte) (byte, bool) {
	if x == 0 {
		return 0, false
	}
	return field.log[x], true
}

// Exp returns 2 to the power of its argument, with multiplication defined
// modulo the irreducible polynomial with which this field was constructed..
func (field *GF256) Exp(x byte) byte {
	return field.exp[x]
}
