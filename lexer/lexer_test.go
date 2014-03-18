package lexer

import (
	"../token"
	"reflect"
	"strings"
	"testing"
)

func LexerTest(t *testing.T) {
	tokChan := Lexer(strings.NewReader(`
package prime

import math

def isPrime(x Int) Bool = {
    isPrime(x, math.Abs(math.Sqrt(x)))
}

def isPrime(x, y Int) Bool = {
    if (y <= 1) {
        true
    } else if (x % y == 0) {
        false
    } else {
        isPrime(x, y - 1)
    }
}

def main = {
    val a = 10
    val d = "a string too"

    println(a)
    println(d)

    println(isPrime(10))
    println(isPrime(7))
    println(isPrime(0))
    println(isPrime(900))
}
`))

	
/*
	package_ := <- tokChan
	switch package_.(type) {
	case token.PackageKeyword:
	default:
		t.Errorf("expected PackageKeyword, found: %s/%s", package_.String(), package_.Value())
	}

	prime := <- tokChan
	switch prime.(type) {
	case token.Ident:
		if prime.String() != "prime" || prime.Value() != "ident" {
			t.Errorf("expected string/value of prime/ident, found: %s/%s", prime.String(), prime.Value())
		}
	default:
		t.Errorf("expected Ident, found: %s/%s", prime.String(), prime.Value())
	}

	import_ := <- tokChan
	switch import_.(type) {
	case token.ImportKeyword:
	default:
		t.Errorf("expected ImportKeyword, found: %s/%s", import_.String(), import_.Value())
	}

	math := <- tokChan
	switch math.(type) {
		
	}
*/

	typeCheck(t, <- tokChan, reflect.TypeOf(token.ImportKeyword{}), func() {})
}

func typeCheck(t *testing.T, tok token.Token, ty reflect.Type, check func()) {
	t.Log(reflect.TypeOf(tok).Name())
	println(ty.Name())
	if reflect.TypeOf(tok).Name() != ty.Name() {
		t.Errorf("wrong type found")
	}
	check()
}
