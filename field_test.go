package gf256

import (
	"testing"
)

func TestField(t *testing.T) {
	testFieldHelper(t, Poly11B, "gf256.New(gf256.Poly11B)", "GF(256/x^8+x^4+x^3+x+1)")
	testFieldHelper(t, Poly11D, "gf256.New(gf256.Poly11D)", "GF(256/x^8+x^4+x^3+x^2+1)")
	testFieldHelper(t, Poly12B, "gf256.New(gf256.Poly12B)", "GF(256/x^8+x^5+x^3+x+1)")
	testFieldHelper(t, Poly12D, "gf256.New(gf256.Poly12D)", "GF(256/x^8+x^5+x^3+x^2+1)")
	testFieldHelper(t, Poly139, "gf256.New(gf256.Poly139)", "GF(256/x^8+x^5+x^4+x^3+1)")
	testFieldHelper(t, Poly13F, "gf256.New(gf256.Poly13F)", "GF(256/x^8+x^5+x^4+x^3+x^2+x+1)")
	testFieldHelper(t, Poly14D, "gf256.New(gf256.Poly14D)", "GF(256/x^8+x^6+x^3+x^2+1)")
	testFieldHelper(t, Poly15F, "gf256.New(gf256.Poly15F)", "GF(256/x^8+x^6+x^4+x^3+x^2+x+1)")
	testFieldHelper(t, Poly163, "gf256.New(gf256.Poly163)", "GF(256/x^8+x^6+x^5+x+1)")
	testFieldHelper(t, Poly165, "gf256.New(gf256.Poly165)", "GF(256/x^8+x^6+x^5+x^2+1)")
	testFieldHelper(t, Poly169, "gf256.New(gf256.Poly169)", "GF(256/x^8+x^6+x^5+x^3+1)")
	testFieldHelper(t, Poly171, "gf256.New(gf256.Poly171)", "GF(256/x^8+x^6+x^5+x^4+1)")
	testFieldHelper(t, Poly177, "gf256.New(gf256.Poly177)", "GF(256/x^8+x^6+x^5+x^4+x^2+x+1)")
	testFieldHelper(t, Poly17B, "gf256.New(gf256.Poly17B)", "GF(256/x^8+x^6+x^5+x^4+x^3+x+1)")
	testFieldHelper(t, Poly187, "gf256.New(gf256.Poly187)", "GF(256/x^8+x^7+x^2+x+1)")
	testFieldHelper(t, Poly18B, "gf256.New(gf256.Poly18B)", "GF(256/x^8+x^7+x^3+x+1)")
	testFieldHelper(t, Poly18D, "gf256.New(gf256.Poly18D)", "GF(256/x^8+x^7+x^3+x^2+1)")
	testFieldHelper(t, Poly19F, "gf256.New(gf256.Poly19F)", "GF(256/x^8+x^7+x^4+x^3+x^2+x+1)")
	testFieldHelper(t, Poly1A3, "gf256.New(gf256.Poly1A3)", "GF(256/x^8+x^7+x^5+x+1)")
	testFieldHelper(t, Poly1A9, "gf256.New(gf256.Poly1A9)", "GF(256/x^8+x^7+x^5+x^3+1)")
	testFieldHelper(t, Poly1B1, "gf256.New(gf256.Poly1B1)", "GF(256/x^8+x^7+x^5+x^4+1)")
	testFieldHelper(t, Poly1BD, "gf256.New(gf256.Poly1BD)", "GF(256/x^8+x^7+x^5+x^4+x^3+x^2+1)")
	testFieldHelper(t, Poly1C3, "gf256.New(gf256.Poly1C3)", "GF(256/x^8+x^7+x^6+x+1)")
	testFieldHelper(t, Poly1CF, "gf256.New(gf256.Poly1CF)", "GF(256/x^8+x^7+x^6+x^3+x^2+x+1)")
	testFieldHelper(t, Poly1D7, "gf256.New(gf256.Poly1D7)", "GF(256/x^8+x^7+x^6+x^4+x^2+x+1)")
	testFieldHelper(t, Poly1DD, "gf256.New(gf256.Poly1DD)", "GF(256/x^8+x^7+x^6+x^4+x^3+x^2+1)")
	testFieldHelper(t, Poly1E7, "gf256.New(gf256.Poly1E7)", "GF(256/x^8+x^7+x^6+x^5+x^2+x+1)")
	testFieldHelper(t, Poly1F3, "gf256.New(gf256.Poly1F3)", "GF(256/x^8+x^7+x^6+x^5+x^4+x+1)")
	testFieldHelper(t, Poly1F5, "gf256.New(gf256.Poly1F5)", "GF(256/x^8+x^7+x^6+x^5+x^4+x^2+1)")
	testFieldHelper(t, Poly1F9, "gf256.New(gf256.Poly1F9)", "GF(256/x^8+x^7+x^6+x^5+x^4+x^3+1)")
}

func testFieldHelper(t *testing.T, p Polynomial, expectGoString string, expectString string) {
	t.Helper()

	type triple struct{ x, y, z byte }

	t.Run(p.String(), func(t *testing.T) {
		field := New(p)

		actualPoly := field.Polynomial()
		actualGoString := field.GoString()
		actualString := field.String()
		if p != actualPoly {
			t.Errorf("Polynomial(): expected %#v, got %#v", p, actualPoly)
		}
		if expectGoString != actualGoString {
			t.Errorf("GoString(): expected %q, got %q", expectGoString, actualGoString)
		}
		if expectString != actualString {
			t.Errorf("String(): expected %q, got %q", expectString, actualString)
		}

		for i := uint(0); i < 256; i++ {
			x := byte(i)
			for j := uint(0); j < 256; j++ {
				y := byte(j)
				testAdd(t, field, x, y)
				testAdd(t, field, y, x)
				testSub(t, field, x, y)
				testSub(t, field, y, x)
				z := slowMul(x, y, p)
				testMul(t, field, x, y, z)
				testMul(t, field, y, x, z)
				testDiv(t, field, z, x, y)
				testDiv(t, field, z, y, x)
			}
		}
		for i := uint(0); i < 256; i++ {
			x := byte(i)
			testMulInv(t, field, x)
		}
		for i := uint(0); i < 256; i++ {
			x := byte(i)
			testExpLog(t, field, x)
			testLogExp(t, field, x)
		}
	})
}

func testAdd(t *testing.T, field *GF256, x byte, y byte) {
	t.Helper()
	z := (x ^ y)
	sum := field.Add(x, y)
	if z != sum {
		t.Errorf("Add(0x%02x, 0x%02x): expected 0x%02x, got 0x%02x", x, y, z, sum)
	}
}

func testSub(t *testing.T, field *GF256, x byte, y byte) {
	t.Helper()
	z := (x ^ y)
	diff := field.Sub(x, y)
	if z != diff {
		t.Errorf("Sub(0x%02x, 0x%02x): expected 0x%02x, got 0x%02x", x, y, z, diff)
	}
}

func testMul(t *testing.T, field *GF256, x byte, y byte, z byte) {
	t.Helper()
	prod := field.Mul(x, y)
	if z != prod {
		t.Errorf("Mul(0x%02x, 0x%02x): expected 0x%02x, got 0x%02x", x, y, z, prod)
	}
}

func testDiv(t *testing.T, field *GF256, z byte, y byte, x byte) {
	t.Helper()
	div, ok := field.Div(z, y)
	switch {
	case y == 0 && ok:
		t.Errorf("Div(0x%02x, 0x%02x): expected divide by zero, got 0x%02x", z, y, div)
	case y == 0:
		return
	case ok && div == x:
		return
	case ok:
		t.Errorf("Div(0x%02x, 0x%02x): expected 0x%02x, got 0x%02x", z, y, x, div)
	default:
		t.Errorf("Div(0x%02x, 0x%02x): expected 0x%02x, got divide by zero", z, y, x)
	}
}

func testMulInv(t *testing.T, field *GF256, x byte) {
	t.Helper()
	y, ok := field.Inv(x)
	switch {
	case x == 0 && ok:
		t.Errorf("Inv(0x%02x): expected divide by zero, got 0x%02x", x, y)
	case x == 0:
		return
	case ok:
		testMul(t, field, x, y, 1)
		testMul(t, field, y, x, 1)
	default:
		t.Errorf("Inv(0x%02x): expected valid element, got divide by zero", x)
	}
}

func testExpLog(t *testing.T, field *GF256, x byte) {
	t.Helper()
	y := field.Exp(x)
	z, ok := field.Log(y)
	switch {
	case ok && x == z:
		// pass
	case ok && x == 0xff && z == 0x00:
		// pass
	case ok:
		t.Errorf("Log(Exp(0x%02x)) aka Log(0x%02x): expected 0x%02x, got 0x%02x", x, y, x, z)
	default:
		t.Errorf("Log(Exp(0x%02x)) aka Log(0x%02x): expected 0x%02x, got failure", x, y, x)
	}
}

func testLogExp(t *testing.T, field *GF256, x byte) {
	t.Helper()

	y, ok := field.Log(x)
	switch {
	case x == 0 && !ok:
		return
	case x == 0:
		t.Errorf("Log(0x%02x): expected failure, got 0x%02x", x, y)
	case !ok:
		t.Errorf("Log(0x%02x): expected valid element, got failure", x)
	default:
		z := field.Exp(y)
		if x != z {
			t.Errorf("Exp(Log(0x%02x)) aka Exp(0x%02x): expected 0x%02x, got 0x%02x", x, y, x, z)
		}
	}
}
