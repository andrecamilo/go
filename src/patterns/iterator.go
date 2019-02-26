package main4

import (
	"errors"
	"log"
)

type Iterator struct {
	tasks    []string
	position int
}

// ErrEOF signals a graceful end of iteration
var ErrEOF = errors.New("EOF")

// Next will return the next task in the slice, or ErrEOF
func (t *Iterator) Next() (int, string, error) {
	t.position++
	if t.position > len(t.tasks) {
		return t.position, "", ErrEOF
	}
	return t.position, t.tasks[t.position-1], nil
}

func main4() {
	tasks := task.GetExamples()
	for {
		i, t, err := tasks.Next()
		if err == task.ErrEOF {
			log.Printf("Done")
			break
		}
		if err != nil {
			log.Fatalf("Unknown error: %s", err)
		}
		log.Printf("Task %d: %s\n", i, t)
	}
}
