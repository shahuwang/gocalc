package main

import ()

type TokenKind int

const (
	BAD_TOKEN TokenKind = iota
	NUMBER_TOKEN
	ADD_OPERATOR_TOKEN
	SUB_OPERATOR_TOKEN
	MUL_OPERATOR_TOKEN
	DIV_OPERATOR_TOKEN
	LEFT_PAREN_TOKEN
	RIGHT_PAREN_TOKEN
	END_OF_LINE_TOKEN
)

type Token struct {
	Kind  TokenKind
	Value float64
	Str   []rune
}

func (t *Token) String() string {
	return string(t.Str)
}

type LexerStatus int

const (
	INITIAL_STATUS LexerStatus = iota
	IN_INT_PART_STATUS
	DOT_STATUS
	IN_FRAC_PART_STATUS
)
