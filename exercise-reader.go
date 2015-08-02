package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (int, error) {
	var i int
	for i = range b {
		b[i] = 'A'
	}
	return i, nil
}

func main() {
	reader.Validate(MyReader{})
}
