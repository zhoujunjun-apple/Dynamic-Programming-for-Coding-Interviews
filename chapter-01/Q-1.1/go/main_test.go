package main

import (
	"fmt"
	"testing"
)

func Test_factorialFunc(t *testing.T) {
	ts := []struct {
		input  uint
		output uint
	}{
		{
			input:  0,
			output: 0,
		},
		{
			input:  1,
			output: 1,
		},
	}

	for _, tt := range ts {
		got := factorialFunc(tt.input)
		if got != tt.output {
			fmt.Printf("input=%d, expected=%d, got=%d\n", tt.input, tt.output, got)
		}
	}
}
