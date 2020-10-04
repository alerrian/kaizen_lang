package token

type TokenType string

const (
	// Special Types
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	LESSTHAN  = "<"
	GRTRTHAN  = ">"
	EQUAL     = "=="
	NOT_EQUAL = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LCURL     = "{"
	RCURL     = "}"
	LBRACK    = "["
	RBRACK    = "]"

	// Keywords
	FUNCTION = "DEF"
	VAR      = "VAR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}
