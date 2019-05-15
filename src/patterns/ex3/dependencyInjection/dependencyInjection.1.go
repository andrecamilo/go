package main

import (
	"io"
	"os"
)

// Creational: Dependency Injection (Setter) pattern
func main() {
	s := NewMyService()
	s.SetWriter(os.Stderr)
	s.WriteHello("world")
}

type MyService struct {
	writer io.Writer
}

func NewMyService() MyService {
	return MyService{}
}
func (s *MyService) SetWriter(w io.Writer) {
	s.writer = w
}

// (Same WriteHello method as last slide)
