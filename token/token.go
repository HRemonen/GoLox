/*
Package token defines the token type and the token struct.
*/
package token

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType   // The type of the toke. See the constants below
	Lexeme  string      // The actual string of the token
	Literal interface{} // The literal value of the token
	Line    int         // Line number where the token was found
}

const (
	// Single-character tokens
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"
	COMMA       = ","
	DOT         = "."
	MINUS       = "-"
	PLUS        = "+"
	SEMICOLON   = ";"
	SLASH       = "/"
	STAR        = "*"

	// One or two character tokens
	BANG          = "!"
	BANG_EQUAL    = "!="
	EQUAL         = "="
	EQUAL_EQUAL   = "=="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="

	// Literals
	IDENTIFIER = "IDENTIFIER"
	STRING     = "STRING"
	NUMBER     = "NUMBER"

	// Keywords
	AND    = "AND"
	CLASS  = "CLASS"
	ELSE   = "ELSE"
	FALSE  = "FALSE"
	FUN    = "FUN"
	FOR    = "FOR"
	IF     = "IF"
	NULL   = "NULL"
	OR     = "OR"
	PRINT  = "PRINT"
	RETURN = "RETURN"
	SUPER  = "SUPER"
	THIS   = "THIS"
	TRUE   = "TRUE"
	VAR    = "VAR"
	WHILE  = "WHILE"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

func New(t TokenType, l string, lit interface{}, line int) *Token {
	return &Token{Type: t, Lexeme: l, Literal: lit, Line: line}
}

func (t *Token) String() string {
	return fmt.Sprintf("%v %v %v", t.Type, t.Lexeme, t.Literal)
}
