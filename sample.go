package main

import (
	"bytes"
	"fmt"
)

var n = 10

func main() {
	out := func1("12345")
	fmt.Println(string(out))
}

func func1(in string) (out []byte) {
	buf := &bytes.Buffer{}
	for i := 0; i < n; i++ {
		buf.WriteString(in)
	}
	out = buf.Bytes()
	return
}
