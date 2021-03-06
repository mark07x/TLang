package lexer

import (
	"testing"

	"github.com/mark07x/TLang/token"
)

func TestNextToken(t *testing.T) {
	input := `
1.0==1.
.2==-2e-04
3!=3
4!=4
5>=5
6>6
7>!7
8<8/
9/=9
10*=10
"Hello World"
"123"
"\n"
"\""
""
'x'
void
true
[][]
_
:
`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.Number, "1.0"},
		{token.Eq, "=="},
		{token.Number, "1."},
		{token.Semicolon, "\n"},

		{token.Number, ".2"},
		{token.Eq, "=="},
		{token.Minus, "-"},
		{token.Number, "2e-04"},
		{token.Semicolon, "\n"},

		{token.Number, "3"},
		{token.NotEq, "!="},
		{token.Number, "3"},
		{token.Semicolon, "\n"},

		{token.Number, "4"},
		{token.NotEq, "!="},
		{token.Number, "4"},
		{token.Semicolon, "\n"},

		{token.Number, "5"},
		{token.GtEq, ">="},
		{token.Number, "5"},
		{token.Semicolon, "\n"},

		{token.Number, "6"},
		{token.Gt, ">"},
		{token.Number, "6"},
		{token.Semicolon, "\n"},

		{token.Number, "7"},
		{token.Gt, ">"},
		{token.Bang, "!"},
		{token.Number, "7"},
		{token.Semicolon, "\n"},

		{token.Number, "8"},
		{token.Lt, "<"},
		{token.Number, "8"},
		{token.Slash, "/"},

		{token.Number, "9"},
		{token.SlashEq, "/="},
		{token.Number, "9"},
		{token.Semicolon, "\n"},

		{token.Number, "10"},
		{token.AsteriskEq, "*="},
		{token.Number, "10"},
		{token.Semicolon, "\n"},

		{token.String, "Hello World"},
		{token.Semicolon, "\n"},
		{token.String, "123"},
		{token.Semicolon, "\n"},
		{token.String, "\\n"},
		{token.Semicolon, "\n"},
		{token.String, "\\\""},
		{token.Semicolon, "\n"},
		{token.String, ""},
		{token.Semicolon, "\n"},

		{token.Character, "x"},
		{token.Semicolon, "\n"},
		{token.Void, "void"},
		{token.Semicolon, "\n"},
		{token.True, "true"},
		{token.Semicolon, "\n"},

		{token.Lbracket, "["},
		{token.Rbracket, "]"},
		{token.Lbracket, "["},
		{token.Rbracket, "]"},
		{token.Semicolon, "\n"},

		{token.Underline, "_"},

		{token.Colon, ":"},

		{token.Eof, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
