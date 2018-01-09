package benchmark

import "strconv"

// 定义获取一个字符
func nextNumString() func() string {
	n := 0
	// closure captures variable n
	return func() string {
		n += 1
		return strconv.Itoa(n)
	}
}
