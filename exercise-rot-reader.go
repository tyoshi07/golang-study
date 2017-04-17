package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (int, error) {
	length, err := reader.r.Read(b)

	for i := 0; i < length; i++ {

		if 'A' <= b[i] && b[i] <= 'Z' {
			b[i] = (((b[i] - 'A') + 13) % 26) + 'A'
		} else if 'a' <= b[i] && b[i] <= 'z' {
			b[i] = (((b[i] - 'a') + 13) % 26) + 'a'
		}
	}
	return length, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
