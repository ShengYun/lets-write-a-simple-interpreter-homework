package main

import "testing"

func TestInterpreter_Expr(t *testing.T) {
	type fields struct {
		lexer        Lexer
		currentToken token
	}
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"test", "1+2", 3},
		{"test", "1+2-3", 0},
		{"test", "1+2  * 3", 7},
		{"test", "1+2* 3  - 8/4", 5},
		{"test", "11+22-3", 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inter := NewInterpreter(tt.input)
			if got := inter.Expr(); got != tt.want {
				t.Errorf("Interpreter.Expr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInterpreter_Expr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		inter := NewInterpreter("1 + 2 * 33 - 33 + 8 / 2")
		inter.Expr()
	}
}
