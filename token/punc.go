package token

type PuncToken struct {}

func (PuncToken) Value() string {
	return "puncuation"
}

type OpenParen struct {
	PuncToken
}

func (OpenParen) String() string {
	return "("
}

type CloseParen struct {
	PuncToken
}

func (CloseParen) String() string {
	return ")"
}

type OpenCurly struct {
	PuncToken
}

func (OpenCurly) String() string {
	return "{"
}

type CloseCurly struct {
	PuncToken
}

func (CloseCurly) String() string {
	return "}"
}

type Assign struct {
	PuncToken
}

func (Assign) String() string {
	return "="
}

type Comma struct {
	PuncToken
}

func (Comma) String() string {
	return ","
}

type Dot struct {
	PuncToken
}

func (Dot) String() string {
	return "."
}

type Mod struct {
	PuncToken
}

func (Mod) String() string {
	return "%"
}

type Minus struct {
	PuncToken
}

func (Minus) String() string {
	return "-"
}

type Equals struct {
	PuncToken
}

func (Equals) String() string {
	return "=="
}

type Lte struct {
	PuncToken
}

func (Lte) String() string {
	return "<="
}

type Lt struct {
	PuncToken
}

func (Lt) String() string {
	return "<"
}
