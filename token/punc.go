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
