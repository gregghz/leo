package parser

import (
	"../lexer"
	"../node"
	"strings"
	"testing"
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
		} else {
			switch pt.Children()[0].(type) {
			case *node.PackageNode:
			default:
				t.Errorf("expected PackageNode")
			}

			switch pt.Children()[1].(type) {
			case *node.ImportNode:
			default:
				t.Errorf("expected ImportNode")
			}

			switch pt.Children()[2].(type) {
			case *node.FuncDef:
				// check children
			default:
				t.Errorf("expected FuncDef")
			}
		}
	default:
		t.Errorf("%#v", pt)
		t.Errorf("found something besides a *ParseTree")
	}
}
