package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	Pos    int
	Status LexerStatus
	Input  []byte
}

func NewLexer(input []byte) *Lexer {
	lexer := new(Lexer)
	lexer.Input = input
	return lexer
}

func (l *Lexer) GetToken() *Token {
	token := new(Token)
	l.Status = INITIAL_STATUS
	for l.Pos < len(l.Input) {
		r, n := utf8.DecodeRune(l.Input[l.Pos:])
		if (l.Status == IN_INT_PART_STATUS || l.Status == IN_FRAC_PART_STATUS) &&
			!unicode.IsDigit(r) && r != '.' {

			token.Kind = NUMBER_TOKEN
			token.Value, _ = strconv.ParseFloat(string(token.Str), 64)
			return token
		}
		if unicode.IsSpace(r) {
			if r == '\n' {
				token.Kind = END_OF_LINE_TOKEN
				return token
			}
			l.Pos = l.Pos + n
			continue
		}
		token.Str = append(token.Str, r)
		l.Pos = l.Pos + n
		switch {
		case r == '+':
			token.Kind = ADD_OPERATOR_TOKEN
			return token
		case r == '-':
			token.Kind = SUB_OPERATOR_TOKEN
			return token
		case r == '*':
			token.Kind = MUL_OPERATOR_TOKEN
			return token
		case r == '/':
			token.Kind = DIV_OPERATOR_TOKEN
			return token
		case r == '(':
			token.Kind = LEFT_PAREN_TOKEN
			return token
		case r == ')':
			token.Kind = RIGHT_PAREN_TOKEN
			return token
		case r == '.':
			if l.Status == IN_INT_PART_STATUS {
				l.Status = DOT_STATUS
			} else {
				fmt.Println("syntax error")
				os.Exit(1)
			}
		case unicode.IsDigit(r):
			if l.Status == INITIAL_STATUS {
				l.Status = IN_INT_PART_STATUS
			} else if l.Status == DOT_STATUS {
				l.Status = IN_FRAC_PART_STATUS
			}
		}
	}
	return token
}
