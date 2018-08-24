package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type tokenType string

const (
	INTEGER tokenType = "INTEGER"
	PLUS    tokenType = "PLUS"
	MINUS   tokenType = "MINUS"
	MUL     tokenType = "MUL"
	DIV     tokenType = "DIV"
	LPAREN  tokenType = "LPAREN"
	RPAREN  tokenType = "RPAREN"
	EOF     tokenType = "EOF"
)

var (
	SyntaxError  = errors.New("Syntax Error")
	InvalidInput = errors.New("Invalid Input")
)

type token struct {
	kind  tokenType
	value string
}

type Lexer struct {
	text        string
	pos         int
	currentChar byte
}

func (lex *Lexer) getNextToken() token {
	for lex.currentChar != '\x00' {
		ch := lex.currentChar

		if isSpace(ch) {
			lex.skipWhitespace()
			continue
		}

		if isDigit(ch) {
			num := lex.integer()
			return token{INTEGER, num}
		}

		if ch == '+' {
			lex.advance()
			return token{PLUS, "+"}
		}

		if ch == '-' {
			lex.advance()
			return token{MINUS, "-"}
		}

		if ch == '*' {
			lex.advance()
			return token{MUL, "*"}
		}

		if ch == '/' {
			lex.advance()
			return token{DIV, "/"}
		}

		if ch == '(' {
			lex.advance()
			return token{LPAREN, "("}
		}

		if ch == ')' {
			lex.advance()
			return token{RPAREN, ")"}
		}

		panic(InvalidInput)
	}

	return token{EOF, "EOF"}
}

func (lex *Lexer) advance() {
	lex.pos++
	if lex.pos > len(lex.text)-1 {
		lex.currentChar = '\x00'
		return
	}
	lex.currentChar = lex.text[lex.pos]
}

func (lex *Lexer) integer() string {
	res := make([]byte, 0)
	for isDigit(lex.currentChar) {
		res = append(res, lex.currentChar)
		lex.advance()
	}
	return string(res)
}

func (lex *Lexer) skipWhitespace() {
	for isSpace(lex.currentChar) {
		lex.advance()
	}
}

func isSpace(ch byte) bool {
	return ch == ' '
}

func isDigit(ch byte) bool {
	return ch >= 0x30 && ch <= 0x39
}

func mustAtoi(s string) int {
	if num, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return num
	}
}

type Interpreter struct {
	lexer        Lexer
	currentToken token
}

func (inter *Interpreter) eat(kind tokenType) {
	if inter.currentToken.kind != kind {
		panic(SyntaxError)
	}
	inter.currentToken = inter.lexer.getNextToken()
}

func (inter *Interpreter) Expr() int {
	res := inter.term()

	for inter.currentToken.kind == PLUS || inter.currentToken.kind == MINUS {
		token := inter.currentToken
		if token.kind == PLUS {
			inter.eat(PLUS)
			res = res + inter.term()
		} else if token.kind == MINUS {
			inter.eat(MINUS)
			res = res - inter.term()
		} else {
			panic(SyntaxError)
		}
	}

	return res
}

func (inter *Interpreter) term() int {
	res := inter.factor()

	for inter.currentToken.kind == MUL || inter.currentToken.kind == DIV {
		token := inter.currentToken
		if token.kind == MUL {
			inter.eat(MUL)
			res = res * inter.factor()
		} else if token.kind == DIV {
			inter.eat(DIV)
			res = res / inter.factor()
		} else {
			panic(SyntaxError)
		}
	}

	return res
}

func (inter *Interpreter) factor() int {
	var res int
	token := inter.currentToken
	if token.kind == INTEGER {
		inter.eat(INTEGER)
		res = mustAtoi(token.value)
	} else if token.kind == LPAREN {
		inter.eat(LPAREN)
		res = inter.Expr()
		inter.eat(RPAREN)
	} else {
		panic(SyntaxError)
	}
	return res
}

func NewInterpreter(input string) Interpreter {
	i := Interpreter{
		lexer: Lexer{
			text:        input,
			pos:         0,
			currentChar: '\x00',
		},
		currentToken: token{},
	}
	i.lexer.currentChar = i.lexer.text[i.lexer.pos]
	i.currentToken = i.lexer.getNextToken()
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i := NewInterpreter(scanner.Text())
		fmt.Printf("%v\n", i.Expr())
	}
}
