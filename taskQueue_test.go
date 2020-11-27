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

	t.Run("removing from empty task queue returns an error", func(t *testing.T) {
		q := TaskQueue{}
		_, err := q.dequeue()

		if err == nil {
			t.Error("dequeue from an empty queue should issue error")
		}
	})

	t.Run("can create a task", func(t *testing.T) {
		task := NewTask("task1", 10)

		if task.ID != "task1" {
			t.Errorf("Created task does not have correct name")
		}
		if task.Data.MaxRuntime != 10 {
			t.Errorf("Created task does not have correct name")
		}
	})

	t.Run("can get last task", func(t *testing.T) {
		q := TaskQueue{}
		task := NewTask("task1", 10)

		q.enqueue(task)

		lastt, err := q.end()
		if err != nil {
			t.Errorf("Error fetching last task from the queue")
		}
		if lastt.ID != "task1" {
			t.Errorf("Wrong task returned as end task")
		}
	})

	t.Run("empty queue returns error when asked for end", func(t *testing.T) {
		q := TaskQueue{}

		_, err := q.end()
		if err == nil {
			t.Errorf("Error empty queue should raise error if asked for last element")
		}
	})
}
