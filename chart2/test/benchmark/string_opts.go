package benchmark

import "strconv"

// nextString is an iterator we use to represent a process
// that returns strings that we want to concatenate in order.
func nextString() func() string {
	n := 0
	// closure captures variable n
	return func() string {
		n += 1
		return strconv.Itoa(n)
	}
}

