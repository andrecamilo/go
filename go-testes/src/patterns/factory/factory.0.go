package main

import (
	"fmt"
	"io"
	"os"
)

func NewWriter(kind string) (io.Writer, error) {
	switch kind {
	case "mywriter":
		return MyWriter{}, nil
	case "stderr":
		return os.Stderr, nil
	default:
		return nil, fmt.Errorf("Kind invalid: %s", kind)
	}
}

type MyWriter struct{}

func (w MyWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("%s", p)
	return len(p), nil
}
