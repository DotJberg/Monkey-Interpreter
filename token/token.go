package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

cont (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//Identifiers and literals
	IDENT = "IDENT" // add, foobar, x, y, ect...
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
