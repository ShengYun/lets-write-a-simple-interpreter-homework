package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	INTEGER = "INTEGER"
	PLUS    = "PLUS"
	MINIUS  = "MINIUS"
	EOF     = "EOF"
)

type token struct {
	kind  string
	value string
}

func (t token) print() {
	fmt.Printf("token(%v, %v)", t.kind, t.value)
}

type Interpreter struct {
	text         string
	pos          int
	currentChar  byte
	currentToken token
}

func NewInterpreter(input string) Interpreter {
	interpreter := Interpreter{
		text: input,
	}
	interpreter.currentChar = interpreter.text[interpreter.pos]
	return interpreter
}

func (interpreter *Interpreter) Expr() int {
	result := interpreter.term()
	for interpreter.currentChar != '\x00' {
		token := interpreter.currentToken
		if token.kind == PLUS {
			interpreter.eat(PLUS)
			result += mustAtoi(interpreter.currentToken.value)
			interpreter.eat(INTEGER)
		}
		if token.kind == MINIUS {
			interpreter.eat(MINIUS)
			result -= mustAtoi(interpreter.currentToken.value)
			interpreter.eat(INTEGER)
		}
	}
	return result
}

func (interpreter *Interpreter) term() int {
	token := interpreter.getNextToken()
	num := mustAtoi(token.value)
	interpreter.eat(INTEGER)
	return num
}

func (interpreter *Interpreter) advance() {
	interpreter.pos++
	if interpreter.pos > len(interpreter.text)-1 {
		interpreter.currentChar = '\x00'
		interpreter.currentToken = token{EOF, "EOF"}
		return
	}
	interpreter.currentChar = interpreter.text[interpreter.pos]
}

func (interpreter *Interpreter) skipWhitespace() {
	for isSpace(interpreter.currentChar) {
		interpreter.advance()
	}
}

func (interpreter *Interpreter) integer() string {
	num := make([]byte, 0)
	for isDigit(interpreter.currentChar) {
		num = append(num, interpreter.currentChar)
		interpreter.advance()
	}
	return string(num)
}

func (interpreter *Interpreter) getNextToken() token {
	for interpreter.currentChar != '\x00' {
		ch := interpreter.currentChar

		if isSpace(ch) {
			interpreter.skipWhitespace()
			continue
		}

		if isDigit(ch) {
			num := interpreter.integer()
			interpreter.currentToken = token{INTEGER, num}
			return interpreter.currentToken
		}

		if ch == '+' {
			interpreter.advance()
			interpreter.currentToken = token{PLUS, "+"}
			return interpreter.currentToken
		}

		if ch == '-' {
			interpreter.advance()
			interpreter.currentToken = token{MINIUS, "-"}
			return interpreter.currentToken
		}
	}

	interpreter.currentToken = token{EOF, "EOF"}
	return interpreter.currentToken
}

func (interpreter *Interpreter) eat(kind string) {
	if interpreter.currentToken.kind == kind {
		interpreter.currentToken = interpreter.getNextToken()
	} else {
		panic("Invalid syntax")
	}
}

func isDigit(ch byte) bool {
	return ch >= 0x30 && ch <= 0x39
}

func isSpace(ch byte) bool {
	return ch == ' '
}

func mustAtoi(s string) int {
	if num, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return num
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i := NewInterpreter(scanner.Text())
		res := i.Expr()
		fmt.Printf("%v\n", res)
	}
}
