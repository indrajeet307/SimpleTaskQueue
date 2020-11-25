package main

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("can create a zero length task queue", func(t *testing.T) {
		q := TaskQueue{}
		fmt.Printf("Queue size is %d\n", q.size())
	})

	t.Run("can add a task to task queue", func(t *testing.T) {
		q := TaskQueue{}
		task := Task{}

		q.enqueue(&task)

		if q.size() != 1 {
			t.Error("Queue length should increase after task is added")
		}

	})

	t.Run("can remove a task from task queue", func(t *testing.T) {
		q := TaskQueue{}
		task := Task{}

		q.enqueue(&task)

		if q.size() != 1 {
			t.Error("Queue length should increase after task is added")
		}

		_, err := q.dequeue()

		if err != nil {
			t.Error("There should be no error when deleting from a populated queue")
		}
		if q.size() != 0 {
			t.Error("Queue length should decrease after task is removed")
		}
	})

	t.Run("removing a from empty task queue returns an error", func(t *testing.T) {
		q := TaskQueue{}
		_, err := q.dequeue()

		if err == nil {
			t.Error("dequeue from an empty queue should issue error")
		}
	})
}
