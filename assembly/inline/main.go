package main

import "github.com/ludwig125/ludwig125_gosample/assembly/inline/hello"

func main() {
	hello.Hello() // inline化する場合-> objdumpの結果からhello呼び出しがなくなる
	// hello.HelloNoInline()
}
