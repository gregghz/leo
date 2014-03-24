package parser

import (
	"../lexer"
	"strings"
	"testing"
	"fmt"
)

func TestParser(t *testing.T) {
	tokChan := lexer.Lexer(strings.NewReader(
		"package prime\n" +
		"import math\n" +
		"def isTrue(b Bool) Bool = true\n",
	))

	pt, errors := Parse(tokChan)

	if len(errors) != 0 {
		t.Errorf("should have no errors")
	}
	
	switch pt.(type) {
	case *ParseTree:
		if len(pt.Children()) != 3 {
			t.Errorf("*ParseTree should have 3 children, found %d", len(pt.Children()))
		}
		a := fmt.Sprintf("%#v\n", pt)
		println(a)
	default:
		t.Errorf("%#v", pt)
		t.Errorf("found something besides a *ParseTree")
	}
}
