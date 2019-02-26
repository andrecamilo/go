package main3

import "os"

func main3() {
	// Get user preference for writer
	kind := "mywriter"
	if len(os.Args) > 1 && os.Args[1] == "stderr" {
		kind = "stderr"
	}
	// Create writer and write some output
	writer, _ := NewWriter(kind)
	writer.Write([]byte("Hello world\n"))
}
