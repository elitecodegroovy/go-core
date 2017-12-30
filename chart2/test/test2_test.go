package test

import (
	"testing"
	"github.com/elitecodegroovy/go-core/util"
)

func TestLongStrSwapCase(t *testing.T){
	s0 := "wE BELIEVE IT'S WORTH TRYING AGAIN WITH A NEW LANGUAGE, " +
		"A CONCURRENT, GARBAGE-COLLECTED LANGUAGE WITH FAST COMPILATION. rEGARDING THE POINTS ABOVE"
	s1 := "We believe it's worth trying again with a new language, a concurrent, " +
		"garbage-collected language with fast compilation. Regarding the points above"
	s2 := util.SwapCase(s1)
	if s0 != s2 {
		t.Errorf("转换后的字符：%s", s2)
	}
}
