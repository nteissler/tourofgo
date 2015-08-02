package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
	"bytes"
)

var ascii_uppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var ascii_lowercase = []byte("abcdefghijklmnopqrstuvwxyz")
var ascii_uppercase_len = len(ascii_uppercase)
var ascii_lowercase_len = len(ascii_lowercase)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
    pos := bytes.IndexByte(ascii_uppercase, b)
    if pos != -1 {
        return ascii_uppercase[(pos+13) % ascii_uppercase_len]
    }
    pos = bytes.IndexByte(ascii_lowercase, b)
    if pos != -1 {
        return ascii_lowercase[(pos+13) % ascii_lowercase_len]
    }
    return b
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	bytesRead, err := rot.r.Read(b)
	//fmt.Println(b)
	var i int = 0
	for ; i < bytesRead; i++ {
		b[i] = rot13(b[i])
	}
	if err != nil {
		return 0, io.EOF
	}
	return i, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
