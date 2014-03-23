package lexer

import (
	"../token"
	"io"
)

var lineNo = 1

func Lexer(in io.Reader) token.TokenChan {
	tc := make(token.TokenChan)

	go func() {
		var b byte
		var err error

		var lex func(byte) (byte, error)
		lex = func(b byte) (byte, error) {
			switch true {
			case isIdentStart(b):
				b, err := ident(string(b), in, tc)
				if err != nil {
					return b, err
				}

				return lex(b)
			case isNumberStart(b):
				b, err := number(string(b), in, tc)
				if err != nil {
					return b, err
				}

				return lex(b)
			case b == '"':
				b, err := str(string(b), in, tc)
				if err != nil {
					return b, err
				}

				return lex(b)
			case b == '(':
				tc <- token.OpenParen{}
			case b == ')':
				tc <- token.CloseParen{}
			case b == '{':
				tc <- token.OpenCurly{}
			case b == '}':
				tc <- token.CloseCurly{}
			case b == '=':
				b, err := equals(string(b), in, tc)
				if err != nil {
					return b, err
				}

				return lex(b)
			case b == ',':
				tc <- token.Comma{}
			case b == '.':
				tc <- token.Dot{}
			case b == '%':
				tc <- token.Mod{}
			case b == '-':
				tc <- token.Minus{}
			case b == '<':
				b, err := lte(string(b), in, tc)
				if err != nil {
					return b, err
				}
			case b == ' ', b == '\t', b == '\n', b == '\r':
				return b, nil
			default:
				tc <- token.NewUnknown("unknown", string(b))
			}
			
			return b, nil
		}
		
		for b, err = nextByte(in); err == nil; b, err = nextByte(in) {
			b, err = lex(b)
		}

		if err == io.EOF {
			tc <- token.Eof{}
		} else {
			panic("lexical error")
		}

		close(tc)
	}()
	
	return tc
}

func isIdentStart(b byte) bool {
	return isAlpha(b) || b == '_'
}

func isNumberStart(b byte) bool {
	return isNumber(b)
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isNumber(b byte) bool {
	return (b >= '0' && b <= '9')
}

func lte(sofar string, in io.Reader, tc token.TokenChan) (byte, error) {
	b, err := nextByte(in)
	if err != nil {
		return b, err
	}

	if b == '=' {
		tc <- token.Lte{}
		return nextByte(in)
	} else {
		tc <- token.Lt{}
		return b, nil
	}
}

func ident(sofar string, in io.Reader, tc token.TokenChan) (byte, error) {
	b, err := nextByte(in)
	if err != nil {
		return b, err
	}

	if !isAlpha(b) && !isNumber(b) && b != '_' {
		sendIdentOrKeyword(sofar, tc)
		return b, err
	} else {
		return ident(sofar + string(b), in, tc)
	}
}

func number(sofar string, in io.Reader, tc token.TokenChan) (byte, error) {
	b, err := nextByte(in)
	if err != nil {
		return b, err
	}

	if !isNumber(b) && b != '.' {
		tc <- token.NewNumber(sofar)
		return b, nil
	} else {
		return number(sofar + string(b), in, tc)
	}
}

func str(sofar string, in io.Reader, tc token.TokenChan) (byte, error) {
	b, err := nextByte(in)
	if err != nil {
		return b, err
	}

	if b == '"' {
		tc <- token.NewString(sofar)
		return nextByte(in)
	} else {
		return str(sofar + string(b), in , tc)
	}
}

func equals(sofar string, in io.Reader, tc token.TokenChan) (byte, error) {
	b, err := nextByte(in)
	if err != nil {
		return b, err
	}

	if b != '=' {
		tc <- token.Assign{}
		return b, nil
	} else {
		tc <- token.Equals{}
		return nextByte(in)
	}
}

func sendIdentOrKeyword(sofar string, tc token.TokenChan) {
	if token.IsKeyword(sofar) {
		tc <- token.StrToKeyword(sofar)
	} else {
		tc <- token.NewIdent(sofar)
	}
}

func nextByte(in io.Reader) (byte, error) {
	b := make([]byte, 1)
	n, err := in.Read(b)

	// we want the error no matter what if n = 0
	if n == 0 {
		return b[0], err
	}

	// we only want to see EOF if n = 0, so we exclude it here
	if err != nil && err != io.EOF {
		return b[0], err
	}

	// if EOF but n != 0, we wait until the next read to get the EOF
	return b[0], nil
}
