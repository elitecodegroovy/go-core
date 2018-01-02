package benchmark

import (
	"testing"
	"fmt"
	"strconv"
)


func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if x := fmt.Sprintf("%d", i); x != strconv.Itoa(i) {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
	b.StopTimer()
}