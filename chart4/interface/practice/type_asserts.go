package practice

import (
	"io"
	"os"
	"fmt"
)

func DoTypeAssert(){
	var r io.Reader
	r = os.Stdout
	f := r.(*os.File) // 成功，因为 f == os.Stdout
	//c := w.(*bytes.Buffer) // panic: 接口是 *os.File, 不是 *bytes.Buffer

	fmt.Printf("f: %v, ", f)
}


func DoTypeAssert2(){
	var r io.Reader
	r = os.Stdout
	rw := r.(io.ReadWriter)//成功: *os.File 有 Read 和 Write 方法
	fmt.Printf("%v, %v", r, rw)

	var c io.Closer
	//panic: interface conversion: interface is nil, not io.ReadWriter [recovered]
	//rw = c.(io.ReadWriter)
	if cw, ok := c.(io.ReadWriter); ok {
		fmt.Printf("%v", cw)
	}else {
		fmt.Printf("\ninterface conversion failed.")
	}

	if c, ok := c.(io.ReadWriter); ok {
		fmt.Printf("%v", c)
	}
}

