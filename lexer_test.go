package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetToken(t *testing.T) {
	data := []string{
		"2 + 5\n",
		"2.5 - 3\n",
		"4*(5-0.9)/12\n",
	}
	for _, item := range data {
		fmt.Println("++++++++++++")
		fmt.Println(item)
		lexer := NewLexer([]byte(item))
		for {
			token := lexer.GetToken()
			if token.Kind == END_OF_LINE_TOKEN {
				break
			}
			fmt.Printf("token is %s\n", token)
		}
		fmt.Println("============")
	}
}

func TestParser(t *testing.T) {
	data := []string{
		"2+5\n",
		"2.5 - 4\n",
		"4*(6-0.9)/12\n",
	}
	for _, item := range data {
		fmt.Println("$$$$$$$$$$$$$$$$$")
		p := NewParser([]byte(item))
		fmt.Printf("%s = %f\n", strings.TrimSpace(item), p.ParseExpression())
		fmt.Println("&&&&&&&&&&&&&&&&")
	}
}
