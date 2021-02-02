package concat

import (
	"bytes"
	"strings"
)

func concat(ss ...string) string {
	var r string
	for _, s := range ss {
		r += s
	}
	return r
}

func concatByteBuffer(ss ...string) string {
	b := bytes.NewBufferString("")
	for _, s := range ss {
		r := strings.NewReader(s)
		r.WriteTo(b)
	}
	return b.String()
}
