package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	ir io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	buffer := make([]byte, len(p))
	n, err = r.ir.Read(buffer)
	for i, v := range buffer {
		lower := strings.ToLower(string(v))[0]
		switch {
		case lower >= 'a' && lower <= 'm':
			p[i] = v + 13
		case lower >= 'n' && lower <= 'z':
			p[i] = v - 13
		default:
			p[i] = v
		}
	}
	return
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
