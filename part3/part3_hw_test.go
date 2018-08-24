package main

import "testing"

func TestInterpreter_Expr(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{"Single digit without space", "1+2", 3},
		{"Multiple digits without space", "12+3", 15},
		{"Single digit with space", "1 + 2", 3},
		{"Single digit with leading space", "  1 + 2", 3},
		{"Single digit with trailing space", "1 - 2  ", -1},
		{"Single digit with leading and trailing space", "  1 + 2  ", 3},
		{"Multiple digits with leading space", "  11 + 2", 13},
		{"Multiple digits with trailing space", "11 - 2  ", 9},
		{"Multiple digits with leading and trailing space", "  11 + 2  ", 13},
		{"long expr", "  1 + 22+33  ", 56},
		{"long expr with minus", "  1 + 22-22+   33  ", 34},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interpreter := NewInterpreter(tt.text)
			if got := interpreter.Expr(); got != tt.want {
				t.Errorf("Interpreter.Expr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInterpreterExpr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		interpreter := NewInterpreter(" 1 + 2 + 3 + 4  - 4 - 6 - 7- 8")
		interpreter.Expr()
	}
}
