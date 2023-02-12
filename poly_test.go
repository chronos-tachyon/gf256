package gf256

import (
	"testing"
)

func TestPolynomial(t *testing.T) {
	testPolynomialHelper(t, Poly11B, "gf256.Poly11B", "x^8+x^4+x^3+x+1", 0x11b)
	testPolynomialHelper(t, Poly11D, "gf256.Poly11D", "x^8+x^4+x^3+x^2+1", 0x11d)
	testPolynomialHelper(t, Poly12B, "gf256.Poly12B", "x^8+x^5+x^3+x+1", 0x12b)
	testPolynomialHelper(t, Poly12D, "gf256.Poly12D", "x^8+x^5+x^3+x^2+1", 0x12d)
	testPolynomialHelper(t, Poly139, "gf256.Poly139", "x^8+x^5+x^4+x^3+1", 0x139)
	testPolynomialHelper(t, Poly13F, "gf256.Poly13F", "x^8+x^5+x^4+x^3+x^2+x+1", 0x13f)
	testPolynomialHelper(t, Poly14D, "gf256.Poly14D", "x^8+x^6+x^3+x^2+1", 0x14d)
	testPolynomialHelper(t, Poly15F, "gf256.Poly15F", "x^8+x^6+x^4+x^3+x^2+x+1", 0x15f)
	testPolynomialHelper(t, Poly163, "gf256.Poly163", "x^8+x^6+x^5+x+1", 0x163)
	testPolynomialHelper(t, Poly165, "gf256.Poly165", "x^8+x^6+x^5+x^2+1", 0x165)
	testPolynomialHelper(t, Poly169, "gf256.Poly169", "x^8+x^6+x^5+x^3+1", 0x169)
	testPolynomialHelper(t, Poly171, "gf256.Poly171", "x^8+x^6+x^5+x^4+1", 0x171)
	testPolynomialHelper(t, Poly177, "gf256.Poly177", "x^8+x^6+x^5+x^4+x^2+x+1", 0x177)
	testPolynomialHelper(t, Poly17B, "gf256.Poly17B", "x^8+x^6+x^5+x^4+x^3+x+1", 0x17b)
	testPolynomialHelper(t, Poly187, "gf256.Poly187", "x^8+x^7+x^2+x+1", 0x187)
	testPolynomialHelper(t, Poly18B, "gf256.Poly18B", "x^8+x^7+x^3+x+1", 0x18b)
	testPolynomialHelper(t, Poly18D, "gf256.Poly18D", "x^8+x^7+x^3+x^2+1", 0x18d)
	testPolynomialHelper(t, Poly19F, "gf256.Poly19F", "x^8+x^7+x^4+x^3+x^2+x+1", 0x19f)
	testPolynomialHelper(t, Poly1A3, "gf256.Poly1A3", "x^8+x^7+x^5+x+1", 0x1a3)
	testPolynomialHelper(t, Poly1A9, "gf256.Poly1A9", "x^8+x^7+x^5+x^3+1", 0x1a9)
	testPolynomialHelper(t, Poly1B1, "gf256.Poly1B1", "x^8+x^7+x^5+x^4+1", 0x1b1)
	testPolynomialHelper(t, Poly1BD, "gf256.Poly1BD", "x^8+x^7+x^5+x^4+x^3+x^2+1", 0x1bd)
	testPolynomialHelper(t, Poly1C3, "gf256.Poly1C3", "x^8+x^7+x^6+x+1", 0x1c3)
	testPolynomialHelper(t, Poly1CF, "gf256.Poly1CF", "x^8+x^7+x^6+x^3+x^2+x+1", 0x1cf)
	testPolynomialHelper(t, Poly1D7, "gf256.Poly1D7", "x^8+x^7+x^6+x^4+x^2+x+1", 0x1d7)
	testPolynomialHelper(t, Poly1DD, "gf256.Poly1DD", "x^8+x^7+x^6+x^4+x^3+x^2+1", 0x1dd)
	testPolynomialHelper(t, Poly1E7, "gf256.Poly1E7", "x^8+x^7+x^6+x^5+x^2+x+1", 0x1e7)
	testPolynomialHelper(t, Poly1F3, "gf256.Poly1F3", "x^8+x^7+x^6+x^5+x^4+x+1", 0x1f3)
	testPolynomialHelper(t, Poly1F5, "gf256.Poly1F5", "x^8+x^7+x^6+x^5+x^4+x^2+1", 0x1f5)
	testPolynomialHelper(t, Poly1F9, "gf256.Poly1F9", "x^8+x^7+x^6+x^5+x^4+x^3+1", 0x1f9)
}

func TestPolynomial_Check(t *testing.T) {
	var panicValue any
	func() {
		defer func() {
			panicValue = recover()
		}()
		Polynomial(NumPolynomials).checkIndex()
	}()
	if panicValue == nil {
		t.Errorf("expected operation to panic, but it did not")
	}
}

func testPolynomialHelper(t *testing.T, p Polynomial, expectGoString string, expectString string, expectValue uint) {
	t.Helper()
	actualGoString := p.GoString()
	actualString := p.String()
	actualValue := p.Value()
	if expectGoString != actualGoString {
		t.Errorf("Polynomial(%d).GoString(): expected %q, got %q", p, expectGoString, actualGoString)
	}
	if expectString != actualString {
		t.Errorf("Polynomial(%d).String(): expected %q, got %q", p, expectString, actualString)
	}
	if expectValue != actualValue {
		t.Errorf("Polynomial(%d).Value(): expected %#x, got %#x", p, expectValue, actualValue)
	}
}
