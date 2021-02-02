package main

import "log"

func main() {
    foo()
    bar()
}

func foo() {
    log.Print("Hello world from foo!")
}

func bar() {
    log.Print("Hello world from bar!")
}

