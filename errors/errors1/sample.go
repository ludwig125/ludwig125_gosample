package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

type temporary interface {
	Temporary() bool
}

func IsTemporary(err error) bool {
	te, ok := errors.Cause(err).(temporary)
	return ok && te.Temporary()
}

type MyError struct {
	Msg    string
	Status bool
}

func (m *MyError) Error() string {
	return fmt.Sprintf("err %s", m.Msg)
}

// Temporary is MyError Status
func (m *MyError) Temporary() bool {
	return m.Status
}

func main() {
	if err := fnA(); err != nil {
		switch cause := errors.Cause(err).(type) {
		case temporary:
			log.Printf("temporary %#v %#v", err.Error(), cause.Temporary())
		default:
			log.Println("not temporary")
		}
		// if IsTemporary(err) {
		// 	log.Println("temporary")
		// 	return
		// }
		// log.Println("not temporary")
	}
}

func fnA() error {
	return &MyError{Msg: "this is error", Status: true}
}
