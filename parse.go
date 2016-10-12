package main

import (
	"fmt"
	"os"
)

type Parser struct {
	ahead_exist bool
	ahead_token *Token
	lex         *Lexer
}

func NewParser(input []byte) *Parser {
	p := new(Parser)
	l := NewLexer(input)
	p.lex = l
	return p
}

func (p *Parser) GetToken() *Token {
	if p.ahead_exist {
		p.ahead_exist = false
		return p.ahead_token
	} else {
		return p.lex.GetToken()
	}
}

func (p *Parser) UnGetToken(token *Token) {
	// 保留预读的token
	p.ahead_exist = true
	p.ahead_token = token
}

func (p *Parser) ParsePrimaryExpression() float64 {
	// 实际上是归约，取单个操作数，比如 2 + 5， 这里取2 和 5
	token := p.GetToken()
	minus_flag := false
	var value float64
	if token.Kind == SUB_OPERATOR_TOKEN {
		minus_flag = true
	} else {
		p.UnGetToken(token)
	}
	token = p.GetToken()
	if token.Kind == NUMBER_TOKEN {
		value = token.Value
	} else if token.Kind == LEFT_PAREN_TOKEN {
		value = p.ParseExpression()
		token = p.GetToken()
		if token.Kind != RIGHT_PAREN_TOKEN {
			fmt.Println("missing ')' error")
			os.Exit(1)
		}
	} else {
		p.UnGetToken(token)
	}
	if minus_flag {
		value = -value
	}
	return value
}

func (p *Parser) ParseTerm() float64 {
	// 归约乘除表达式为一个数
	value1 := p.ParsePrimaryExpression()
	for {
		token := p.GetToken()
		if token.Kind != MUL_OPERATOR_TOKEN && token.Kind != DIV_OPERATOR_TOKEN {
			p.UnGetToken(token)
			break
		}
		value2 := p.ParsePrimaryExpression()
		if token.Kind == MUL_OPERATOR_TOKEN {
			value1 *= value2
		} else if token.Kind == DIV_OPERATOR_TOKEN {
			value1 /= value2
		}
	}
	return value1
}

func (p *Parser) ParseExpression() float64 {
	// 抽象的简化，把所有的表达式抽象成一个加法或者一个减法
	// 先获取操作符左边的数, 有加号或减号，获取操作符右边的数，加减，返回值
	value1 := p.ParseTerm()
	for {
		token := p.GetToken()
		if token.Kind != ADD_OPERATOR_TOKEN && token.Kind != SUB_OPERATOR_TOKEN {
			p.UnGetToken(token)
			break
		}
		value2 := p.ParseTerm()
		if token.Kind == ADD_OPERATOR_TOKEN {
			value1 += value2
		} else if token.Kind == SUB_OPERATOR_TOKEN {
			value1 -= value2
		} else {
			p.UnGetToken(token)
		}
	}
	return value1
}
