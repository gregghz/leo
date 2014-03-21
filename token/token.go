package token

type TokenChan chan Token

var keywords = map[string]Token {
	"package": PackageKeyword{},
	"import": ImportKeyword{},
	"def": DefKeyword{},
	"if": IfKeyword{},
	"else": ElseKeyword{},
	"val": ValKeyword{},
	"lazy": LazyKeyword{},
	"true": TrueKeyword{},
	"false": FalseKeyword{},
}

func IsKeyword(sofar string) bool {
	for kw, _ := range keywords {
		if kw == sofar {
			return true
		}
	}

	return false
}

func StrToKeyword(sofar string) Token {
	return keywords[sofar]
}

type Token interface {
	String() string
	Value() string
}

type Ident struct {
	name, value string
}

func NewIdent(name string) Ident {
	return Ident{name, "ident"}
}

func (i Ident) String() string {
	return i.name
}

func (i Ident) Value() string {
	return i.value
}

type Number struct {
	name, value string
}

func NewNumber(value string) Number {
	return Number{"number", value}
}

func (n Number) String() string {
	return n.name
}

func (n Number) Value() string {
	return n.value
}

type Eof struct {
	KeywordToken
}

func (Eof) String() string {
	return "EOF"
}

