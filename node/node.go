package node

import (
	"../token"
	"fmt"
	"reflect"
)

func Eq(a, b Node) bool {
	return reflect.DeepEqual(a, b)
}

type Node interface {
	Children() []Node
	SetChildren([]Node)
	Go() string
}

type PackageNode struct {
	name string
}

func NewPackageNode(ident token.Token) *PackageNode {
	return &PackageNode{ident.String()}
}

func (PackageNode) Children() []Node {
	return []Node{}
}

func (PackageNode) SetChildren(nodes []Node) {
	
}

func (p PackageNode) Go() string {
	return fmt.Sprintf("package %s", p.name)
}

type ImportNode struct {
	name string
}

func NewImportNode(ident token.Token) *ImportNode {
	return &ImportNode{ident.String()}
}

func (ImportNode) Children() []Node {
	return []Node{}
}

func (ImportNode) SetChildren(nodes []Node) {
	
}

func (i ImportNode) Go() string {
	return fmt.Sprintf("import \"%s\"", i.name)
}

type ParamListNode struct {
	children []Node
}

func NewParamListNode(children []Node) *ParamListNode {
	return &ParamListNode{children}
}

func (p ParamListNode) Children() []Node {
	return p.children
}

func (p *ParamListNode) SetChildren(nodes []Node) {
	p.children = nodes
}

func (p ParamListNode) Go() string {
	str := "("
	for i, param := range p.children {
		if i != 0 {
			str += ","
		}

		switch param.(type) {
		case TypedParam:
			str += fmt.Sprintf("%s %s", param.(TypedParam).name, param.(TypedParam).type_)
		}
	}
	return str + ")"
}

type TypedParam struct {
	name, type_ string
}

func NewTypedParam(name, type_ token.Token) *TypedParam {
	return &TypedParam{name.String(), type_.String()}
}

func (TypedParam) Children() []Node {
	return []Node{}
}

func (TypedParam) SetChildren([]Node) {
	
}

func (TypedParam) Go() string {
	return ""
}

type UntypedParam struct {
	name string
}

func NewUntypedParam(name token.Token) *UntypedParam {
	return &UntypedParam{name.String()}
}

func (UntypedParam) Children() []Node {
	return []Node{}
}

func (UntypedParam) SetChildren([]Node) {
	
}

func (UntypedParam) Go() string {
	return ""
}

type ExprNode struct {

}

func NewExprNode() ExprNode {
	return ExprNode{}
}

func (ExprNode) Children() []Node {
	return []Node{}
}

func (ExprNode) SetChildren([]Node) {

}

func (ExprNode) Go() string {
	return ""
}

type FuncDef struct {
	name string
	children []Node
}

func NewFuncDef(ident token.Token, params, expr Node) *FuncDef {
	return &FuncDef{
		ident.String(),
		[]Node{params, expr},
	}
}

func (f FuncDef) Children() []Node {
	return f.children
}

func (f *FuncDef) SetChildren(nodes []Node) {
	f.children = nodes
}

func (f FuncDef) Go() string {
	str := "func " + f.name + f.children[0].Go()
	str += "{"
	str += f.children[1].Go()
	str += "}"
	return str
}
