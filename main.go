package main

import "fmt"

type Reader struct {
	r int
}

type Writer struct {
	w int
}

type ReadWriter struct {
	*Reader
	*Writer
}

func main() {
	var r = Reader{}
	var w = Writer{}
	var rw = ReadWriter{
		Reader: &r,
		Writer: &w,
	}
	fmt.Println(rw)
}
