package test

import (
	"github.com/elitecodegroovy/go-core/util"
	"testing"
)

func TestMultipleSwapCase(t *testing.T) {
	var ts = []struct {
		input    string
		expected string
	}{
		{"abc", "ABC"},
		{"Abc", "aBC"},
		{"abC", "ABc"},
		{"a★★", "A★★"},
		{"a好bc", "A好BC"},
		{"こんにちは,Miss Li", "こんにちは,mISS lI"},
		{"こんにちは！", "こんにちは！"},
	}
	for _, in := range ts {
		if got := util.SwapCase(in.input); in.expected != got {
			t.Errorf("expected :(%q) , real := %v", in.expected, got)
		}
	}
}
