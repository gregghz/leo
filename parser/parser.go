package parser

import (
	"fmt"
)

import(
	"../node"
	"../token"
)

type ParseTree struct {
	children []node.Node
}

func (p ParseTree) Children() []node.Node {
	return p.children
}

func (p *ParseTree) SetChildren(nodes []node.Node) {
	p.children = nodes
}

func (p ParseTree) Go() string {
	str := ""
	for _, node := range p.children {
		str += fmt.Sprintf("%s\n", node.Go())
	}
	return str
}

func Parse(tc token.TokenChan) (node.Node, []error) {
	errors := []error{}

	pt, tok, err := package_(<-tc, tc, &ParseTree{})
	if err != nil {
		errors = append(errors, err)
	}

	pt, tok, err = zeroOrMore(import_, tok, tc, pt)
	if err != nil {
		errors = append(errors, err)
	}

	pt, tok, err = zeroOrMore(def, tok, tc, pt)
	if err != nil {
		errors = append(errors, err)
	}

	return pt, errors
}

type parseFunc func(token.Token, token.TokenChan, node.Node) (node.Node, token.Token, error)

func import_(tok token.Token, tc token.TokenChan, pt node.Node) (node.Node, token.Token, error) {
	err := requireKeyword("import", tok)
	if err != nil {
		return pt, tok, err
	}

	ident, err := requireIdent(<-tc)
	if err != nil {
		return pt, ident, err
	}

	pt.SetChildren(append(pt.Children(), node.NewImportNode(ident)))
	
	return pt, <-tc, nil
}

func package_(tok token.Token, tc token.TokenChan, pt node.Node) (node.Node, token.Token, error) {
	err := requireKeyword("package", tok)
	if err != nil {
		return pt, tok, err
	}

	ident, err := requireIdent(<-tc)
	if err != nil {
		return pt, ident, err
	}

	pt.SetChildren(append(pt.Children(), node.NewPackageNode(ident)))

	return pt, <-tc, nil
}

func def(tok token.Token, tc token.TokenChan, pt node.Node) (node.Node, token.Token, error) {
	err := requireKeyword("def", tok)
	if err != nil {
		return pt, tok, err
	}

	ident, err := requireIdent(<-tc)
	params, tok, err := zeroOrOne(paramList, <-tc, tc, &node.ParamListNode{})
	expr, tok, err := requireExpr(tok, tc, &node.ExprNode{})

	pt.SetChildren(append(pt.Children(), node.NewFuncDef(ident, params, expr)))
	
	return pt, tok, nil
}

func requireExpr(tok token.Token, tc token.TokenChan, exprNode node.Node) (node.Node, token.Token, error) {
	return exprNode, tok, nil
}

func paramList(tok token.Token, tc token.TokenChan, paramListNode node.Node) (node.Node, token.Token, error) {
	err := requireOpenParen(tok)
	if err != nil {
		return paramListNode, tok, err
	}

	ident, hasList := optionalIdent(<-tc)
	if hasList {
		var nextToken token.Token

		type_, hasType := optionalIdent(<-tc)
		if hasType {
			paramListNode.SetChildren(append(paramListNode.Children(), node.NewTypedParam(ident, type_)))
			nextToken = <-tc
		} else {
			paramListNode.SetChildren(append(paramListNode.Children(), node.NewUntypedParam(ident)))
			nextToken = type_
		}
		
		for keepGoing := true; keepGoing; {
			switch nextToken.(type) {
			case token.Comma:
				println("comma")
				ident, err := requireIdent(<-tc)
				if err != nil {
					return paramListNode, tok, err
				}
				
				type_, hasType := optionalIdent(<-tc)
				if hasType {
					paramListNode.SetChildren(append(paramListNode.Children(), node.NewTypedParam(ident, type_)))
					nextToken = <-tc
				} else {
					paramListNode.SetChildren(append(paramListNode.Children(), node.NewUntypedParam(ident)))
					nextToken = type_
				}
			default:
				keepGoing = false
			}
		}

		err = requireCloseParen(<-tc)
	} else {
		err = requireCloseParen(ident)
	}

	if err != nil {
		return paramListNode, tok, err
	}

	return paramListNode, <-tc, nil
}

func optionalIdent(tok token.Token) (token.Token, bool) {
	switch tok.(type) {
	case token.Ident:
		return tok, true
	default:
		return tok, false
	}
}

func requireKeyword(name string, tok token.Token) error {
	if name != tok.String() {
		return fmt.Errorf("expected %s, found %s", name, tok.String())
	}
	return nil
}

func requireIdent(tok token.Token) (token.Token, error) {
	switch tok.(type) {
	case token.Ident:
	default:
		return tok, fmt.Errorf("expected an identifier, found %s", tok.String())
	}
	return tok, nil
}

func requireOpenParen(tok token.Token) error {
	switch tok.(type) {
	case token.OpenParen:
	default:
		return fmt.Errorf("expected '(', found %s", tok.String())
	}
	return nil
}

func requireCloseParen(tok token.Token) error {
	switch tok.(type) {
	case token.CloseParen:
	default:
		return fmt.Errorf("expected ')', found %s", tok.String())
	}
	return nil
}

func zeroOrOne(rule parseFunc, tok token.Token, tc token.TokenChan, pt node.Node) (node.Node, token.Token, error) {
	pt, tok, _ = rule(tok, tc, pt)
	return pt, tok, nil
}

func zeroOrMore(rule parseFunc, tok token.Token, tc token.TokenChan, pt node.Node) (node.Node, token.Token, error) {
	pt, tok, err := rule(tok, tc, pt)
	for err == nil {
		pt, tok, err = rule(tok, tc, pt)
	}

	return pt, tok, nil
}
