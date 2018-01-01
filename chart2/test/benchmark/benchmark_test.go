package benchmark

import (
	"testing"
	"bytes"
)
var global string

func benchmarkNBufferString(b *testing.B, numConcat int) {
	var ns string
	for i := 0; i < b.N; i++ {
		next := nextString()
		buffer := bytes.NewBufferString("")
		for u := 0; u < numConcat; u++ {
			buffer.WriteString(next())
		}
		ns = buffer.String()
	}
	global = ns
}

func BenchmarkBufferString(b *testing.B) {
	benchmarkNBufferString(b, 10)
}
