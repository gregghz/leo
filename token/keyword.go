package token

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
