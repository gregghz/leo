package lexer

import (
	"../token"
	"strings"
	"testing"
)

func TestNumberLexing(t *testing.T) {
	tokChan := Lexer(strings.NewReader("(123)"))

	openParen := <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	number := <-tokChan
	switch number.(type) {
	case token.Number:
		if number.Value() != "123" {
			t.Errorf("expected 123")
		}
	default:
		t.Errorf("wrong token type: %s/%s", number.String(), number.Value())
	}

	closeParen := <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

}



func TestLexer(t *testing.T) {
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

	// package prime
	package_ := <-tokChan
	switch package_.(type) {
	case token.PackageKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", package_.String(), package_.Value())
	}

	prime := <-tokChan
	switch prime.(type) {
	case token.Ident:
		if prime.String() != "prime" {
			t.Errorf("wrong identifier: %s", prime.String())
		}
	default:
		t.Errorf("wrong token type: %s/%s", prime.String(), prime.Value())
	}

	// import math
	import_ := <-tokChan
	switch import_.(type) {
	case token.ImportKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", import_.String(), import_.Value())
	}

	math := <-tokChan
	switch math.(type) {
	case token.Ident:
		if math.String() != "math" {
			t.Errorf("wrong identifier: %s", math.String())
		}
	default:
		t.Errorf("wrong token type: %s/%s", math.String(), math.Value())
	}

	// def isPrime(x Int) Bool = {
	def := <-tokChan
	switch def.(type) {
	case token.DefKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", def.String(), def.Value())
	}

	isPrime := <-tokChan
	switch isPrime.(type) {
	case token.Ident:
		if isPrime.String() != "isPrime" {
			t.Errorf("wrong identifier: %s", isPrime.String())
		}
	default:
		t.Errorf("wrong token type: %s/%s", isPrime.String(), isPrime.Value())
	}

	openParen := <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	x := <-tokChan
	switch x.(type) {
	case token.Ident:
		if x.String() != "x" {
			t.Errorf("wrong identifier: %s", x.String())
		}
	default:
		t.Errorf("wrong token type: %s/%s", x.String(), x.Value())
	}

	Int := <-tokChan
	switch Int.(type) {
	case token.Ident:
		if Int.String() != "Int" {
			t.Errorf("wrong identifier: %s", Int.String())
		}
	default:
		t.Errorf("wrong token type: %s/%s", Int.String(), Int.Value())
	}

	closeParen := <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:

	}

	Bool := <-tokChan
	switch Bool.(type) {
	case token.Ident:
		if Bool.String() != "Bool" {
			t.Errorf("wrong identifier: %s", Int.String())
		}
	default:
		t.Errorf("wrong token type: %s/%s", Bool.String(), Bool.Value())
	}

	assign := <-tokChan
	switch assign.(type) {
	case token.Assign:
	default:
		t.Errorf("wrong token type: %s/%s", assign.String(), assign.Value())
	}

	openCurly := <-tokChan
	switch openCurly.(type) {
	case token.OpenCurly:
	default:
		t.Errorf("wrong token type")
	}

	// isPrime(x, math.Abs(math.Sqrt(x)))
	isPrime = <-tokChan
	switch isPrime.(type) {
	case token.Ident:
		if isPrime.String() != "isPrime" {
			t.Errorf("wrong identifier: %s", isPrime.String())
		}
	default:
		t.Errorf("wrong token type: %s/%s", isPrime.String(), isPrime.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	x = <-tokChan
	switch x.(type) {
	case token.Ident:
		if x.String() != "x" {
			t.Errorf("wrong identifier: %s", x.String())
		}
	default:
		t.Errorf("wrong token type: %s/%s", x.String(), x.Value())
	}

	comma := <-tokChan
	switch comma.(type) {
	case token.Comma:
	default:
		t.Errorf("wrong token type: %s/%s", comma.String(), comma.Value())
	}

	math = <-tokChan
	switch math.(type) {
	case token.Ident:
		if math.String() != "math" {
			t.Errorf("wrong identifier: %s", math.String())
		}
	default:
		t.Errorf("wrong token type: %s/%s", math.String(), math.Value())
	}

	dot := <-tokChan
	switch dot.(type) {
	case token.Dot:
	default:
		t.Errorf("wrong token type: %s/%s", dot.String(), dot.Value())
	}

	abs := <-tokChan
	switch abs.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", abs.String(), abs.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	math = <-tokChan
	switch math.(type) {
	case token.Ident:
		if math.String() != "math" {
			t.Errorf("wrong identifier: %s", math.String())
		}
	default:
		t.Errorf("wrong token type: %s/%s", math.String(), math.Value())
	}

	dot = <-tokChan
	switch dot.(type) {
	case token.Dot:
	default:
		t.Errorf("wrong token type: %s/%s", dot.String(), dot.Value())
	}

	sqrt := <-tokChan
	switch sqrt.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", sqrt.String(), sqrt.Value())
	}
	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	x = <-tokChan
	switch x.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", x.String(), x.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	// }
	closeCurly := <-tokChan
	switch closeCurly.(type) {
	case token.CloseCurly:
	default:
		t.Errorf("wrong token type: %s/%s", closeCurly.String(), closeCurly.Value())
	}

	// def isPrime(x, y Int) Bool = {
	def = <-tokChan
	switch def.(type) {
	case token.DefKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", def.String(), def.Value())
	}

	isPrime = <-tokChan
	switch isPrime.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", isPrime.String(), isPrime.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	x = <-tokChan
	switch x.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", x.String(), x.Value())
	}

	comma = <-tokChan
	switch comma.(type) {
	case token.Comma:
	default:
		t.Errorf("wrong token type: %s/%s", comma.String(), comma.Value())
	}

	y := <-tokChan
	switch y.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", y.String(), y.Value())
	}

	Int = <-tokChan
	switch Int.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", Int.String(), Int.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	Bool = <-tokChan
	switch Bool.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", Bool.String(), Bool.Value())
	}

	assign = <-tokChan
	switch assign.(type) {
	case token.Assign:
	default:
		t.Errorf("wrong token type: %s/%s", assign.String(), assign.Value())
	}

	openCurly = <-tokChan
	switch openCurly.(type) {
	case token.OpenCurly:
	default:
		t.Errorf("wrong token type: %s/%s", openCurly.String(), openCurly.Value())
	}

	// if (y <= 1) {
	if_ := <-tokChan
	switch if_.(type) {
	case token.IfKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", if_.String(), if_.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	y = <-tokChan
	switch y.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", y.String(), y.Value())
	}

	lte := <-tokChan
	switch lte.(type) {
	case token.Lte:
	default:
		t.Errorf("wrong token type: %s/%s", lte.String(), lte.Value())
	}

	one := <-tokChan
	switch one.(type) {
	case token.Number:
	default:
		t.Errorf("wrong token type: %s/%s", one.String(), one.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	openCurly = <-tokChan
	switch openCurly.(type) {
	case token.OpenCurly:
	default:
		t.Errorf("wrong token type: %s/%s", openCurly.String(), openCurly.Value())
	}

	// true
	true_ := <-tokChan
	switch true_.(type) {
	case token.TrueKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", true_.String(), true_.Value())
	}

	// } else if (x % y == 0) {
	closeCurly = <-tokChan
	switch closeCurly.(type) {
	case token.CloseCurly:
	default:
		t.Errorf("wrong token type: %s/%s", closeCurly.String(), closeCurly.Value())
	}

	else_ := <-tokChan
	switch else_.(type) {
	case token.ElseKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", else_.String(), else_.Value())
	}

	if_ = <-tokChan
	switch if_.(type) {
	case token.IfKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", if_.String(), if_.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	x = <-tokChan
	switch x.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", x.String(), x.Value())
	}

	mod := <-tokChan
	switch mod.(type) {
	case token.Mod:
	default:
		t.Errorf("wrong token type: %s/%s", mod.String(), mod.Value())
	}

	y = <-tokChan
	switch y.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", y.String(), y.Value())
	}

	equals := <-tokChan
	switch equals.(type) {
	case token.Equals:
	default:
		t.Errorf("wrong token type: %s/%s", equals.String(), equals.Value())
	}

	number := <-tokChan
	switch number.(type) {
	case token.Number:
	default:
		t.Errorf("wrong token type: %s/%s", number.String(), number.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	openCurly = <-tokChan
	switch openCurly.(type) {
	case token.OpenCurly:
	default:
		t.Errorf("wrong token type: %s/%s", openCurly.String(), openCurly.Value())
	}

	// false
	false_ := <-tokChan
	switch false_.(type) {
	case token.FalseKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", false_.String(), false_.Value())
	}

	// } else {
	closeCurly = <-tokChan
	switch closeCurly.(type) {
	case token.CloseCurly:
	default:
		t.Errorf("wrong token type: %s/%s", closeCurly.String(), closeCurly.Value())
	}

	else_ = <-tokChan
	switch else_.(type) {
	case token.ElseKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", else_.String(), else_.Value())
	}

	openCurly = <-tokChan
	switch openCurly.(type) {
	case token.OpenCurly:
	default:
		t.Errorf("wrong token type: %s/%s", openCurly.String(), openCurly.Value())
	}

	// isPrime(x, y - 1)
	isPrime = <-tokChan
	switch isPrime.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", isPrime.String(), isPrime.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	x = <-tokChan
	switch x.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", x.String(), x.Value())
	}

	comma = <-tokChan
	switch comma.(type) {
	case token.Comma:
	default:
		t.Errorf("wrong token type: %s/%s", comma.String(), comma.Value())
	}

	y = <-tokChan
	switch y.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", y.String(), y.Value())
	}

	minus := <-tokChan
	switch minus.(type) {
	case token.Minus:
	default:
		t.Errorf("wrong token type: %s/%s", minus.String(), minus.Value())
	}

	number = <-tokChan
	switch number.(type) {
	case token.Number:
	default:
		t.Errorf("wrong token type: %s/%s", number.String(), number.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	// }
	closeCurly = <-tokChan
	switch closeCurly.(type) {
	case token.CloseCurly:
	default:
		t.Errorf("wrong token type: %s/%s", closeCurly.String(), closeCurly.Value())
	}

	// }
	closeCurly = <-tokChan
	switch closeCurly.(type) {
	case token.CloseCurly:
	default:
		t.Errorf("wrong token type: %s/%s", closeCurly.String(), closeCurly.Value())
	}

	// def main = {
	def = <-tokChan
	switch def.(type) {
	case token.DefKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", def.String(), def.Value())
	}

	main := <-tokChan
	switch main.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", main.String(), main.Value())
	}

	assign = <-tokChan
	switch assign.(type) {
	case token.Assign:
	default:
		t.Errorf("wrong token type: %s/%s", assign.String(), assign.Value())
	}

	openCurly = <-tokChan
	switch openCurly.(type) {
	case token.OpenCurly:
	default:
		t.Errorf("wrong token type: %s/%s", openCurly.String(), openCurly.Value())
	}

	// val a = 10
	val := <-tokChan
	switch val.(type) {
	case token.ValKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", val.String(), val.Value())
	}

	a := <-tokChan
	switch a.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", a.String(), a.Value())
	}

	assign = <-tokChan
	switch assign.(type) {
	case token.Assign:
	default:
		t.Errorf("wrong token type: %s/%s", assign.String(), assign.Value())
	}

	number = <-tokChan
	switch number.(type) {
	case token.Number:
	default:
		t.Errorf("wrong token type: %s/%s", number.String(), number.Value())
	}

	// val d = "a string too"
	val = <-tokChan
	switch val.(type) {
	case token.ValKeyword:
	default:
		t.Errorf("wrong token type: %s/%s", val.String(), val.Value())
	}

	d := <-tokChan
	switch d.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", d.String(), d.Value())
	}

	assign = <-tokChan
	switch assign.(type) {
	case token.Assign:
	default:
		t.Errorf("wrong token type: %s/%s", assign.String(), assign.Value())
	}

	string := <-tokChan
	switch string.(type) {
	case token.String:
	default:
		t.Errorf("wrong token type: %s/%s", string.String(), string.Value())
	}

	// println(a)
	println_ := <-tokChan
	switch println_.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", println_.String(), println_.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	a = <-tokChan
	switch a.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", a.String(), a.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	// println(d)
	println_ = <-tokChan
	switch println_.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", println_.String(), println_.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	a = <-tokChan
	switch a.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", a.String(), a.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	// println(isPrime(10))
	println_ = <-tokChan
	switch println_.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", println_.String(), println_.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	isPrime = <-tokChan
	switch isPrime.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", isPrime.String(), isPrime.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	number = <-tokChan
	switch number.(type) {
	case token.Number:
	default:
		t.Errorf("wrong token type: %s/%s", number.String(), number.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	// println(isPrime(7))
	println_ = <-tokChan
	switch println_.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", println_.String(), println_.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	isPrime = <-tokChan
	switch isPrime.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", isPrime.String(), isPrime.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	number = <-tokChan
	switch number.(type) {
	case token.Number:
	default:
		t.Errorf("wrong token type: %s/%s", number.String(), number.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}
	
	// println(isPrime(0))
	println_ = <-tokChan
	switch println_.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", println_.String(), println_.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	isPrime = <-tokChan
	switch isPrime.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", isPrime.String(), isPrime.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	number = <-tokChan
	switch number.(type) {
	case token.Number:
	default:
		t.Errorf("wrong token type: %s/%s", number.String(), number.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}
	
	// println(isPrime(900))
	println_ = <-tokChan
	switch println_.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", println_.String(), println_.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	isPrime = <-tokChan
	switch isPrime.(type) {
	case token.Ident:
	default:
		t.Errorf("wrong token type: %s/%s", isPrime.String(), isPrime.Value())
	}

	openParen = <-tokChan
	switch openParen.(type) {
	case token.OpenParen:
	default:
		t.Errorf("wrong token type: %s/%s", openParen.String(), openParen.Value())
	}

	number = <-tokChan
	switch number.(type) {
	case token.Number:
	default:
		t.Errorf("wrong token type: %s/%s", number.String(), number.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}

	closeParen = <-tokChan
	switch closeParen.(type) {
	case token.CloseParen:
	default:
		t.Errorf("wrong token type: %s/%s", closeParen.String(), closeParen.Value())
	}
	
	// }
	closeCurly = <-tokChan
	switch closeCurly.(type) {
	case token.CloseCurly:
	default:
		t.Errorf("wrong token type: %s/%s", closeCurly.String(), closeCurly.Value())
	}
}
