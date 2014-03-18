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

type KeywordToken struct {}

func (KeywordToken) Value() string {
	return "keyword"
}


type PackageKeyword struct {
	KeywordToken
}

func (PackageKeyword) String() string {
	return "package"
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


type ImportKeyword struct {
	KeywordToken
}

func (ImportKeyword) String() string {
	return "import"
}


type DefKeyword struct {
	KeywordToken
}

func (DefKeyword) String() string {
	return "def"
}


type IfKeyword struct {
	KeywordToken
}

func (IfKeyword) String() string {
	return "if"
}


type ElseKeyword struct {
	KeywordToken
}

func (ElseKeyword) String() string {
	return "else"
}


type ValKeyword struct {
	KeywordToken
}

func (ValKeyword) String() string {
	return "val"
}


type LazyKeyword struct {
	KeywordToken
}

func (LazyKeyword) String() string {
	return "lazy"
}


type TrueKeyword struct {
	KeywordToken
}

func (TrueKeyword) String() string {
	return "true"
}


type FalseKeyword struct {
	KeywordToken
}

func (FalseKeyword) String() string {
	return "false"
}


type Eof struct {
	KeywordToken
}

func (Eof) String() string {
	return "EOF"
}

