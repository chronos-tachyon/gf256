// Package gf256 implements arithmetic over Galois Fields with characteristic 2
// and exponent 8.  Such fields are frequently written as "GF(256)" or
// "GF(2**8)".
//
// Many algorithms in error correction and cryptography are defined in terms of
// GF(256) itself or in terms of polynomial fields with coefficients in
// GF(256).  Indeed, GF(2**k) for k > 1 is itself a polynomial field with
// coefficients taken from GF(2), where GF(2) is the finite field consisting of
// a single bit.
//
// While addition and subtraction are always defined in the same way for all
// instances of GF(256), multiplication and division must be defined modulo
// some irreducible polynomial with coefficients in GF(2).  There are 16 such
// polynomials, and therefore there are 16 different ways to instantiate
// GF(256).  Poly11D, which represents the polynomial x**8 + x**4 + x**3 + x**2
// + 1, is a popular choice, but there are few reasons to select one polynomial
// over another except for compatibility with existing algorithms.
package gf256
