package main

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Test can create a zero length task queue", func(t *testing.T){
		q := TaskQueue {}
		fmt.Printf("Queue size is %d\n", q.size())
	})
}

