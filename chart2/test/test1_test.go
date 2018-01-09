package test

import (
	"github.com/elitecodegroovy/go-core/util"
	"testing"
)

func TestReverse(t *testing.T) {
	s := "123456789"
	sReverse := util.Reverse(s)
	if "987654321" != sReverse {
		t.Errorf("返回字符串: %s", sReverse)
	}
}

func TestSwapCase(t *testing.T) {
	s1 := "testSwapCase"
	s2 := util.SwapCase(s1)
	if "TESTsWAPcASE" != s2 {
		t.Errorf("返回字符串: %s", s2)
	}

}
