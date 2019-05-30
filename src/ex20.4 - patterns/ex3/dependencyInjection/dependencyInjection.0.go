package main

import (
	"fmt"
	"io"
	"os"
)

// Creational: Dependency Injection (Constructor) pattern

func main() {
	s := NewMyService(os.Stderr)
	s.WriteHello("world")
}

type MyService struct {
	writer io.Writer
}

func NewMyService(writer io.Writer) MyService {
	return MyService{
		writer: writer,
	}
}
func (s *MyService) WriteHello(m string) {
	fmt.Fprintf(s.writer, "Hello %s\n", m)
}
