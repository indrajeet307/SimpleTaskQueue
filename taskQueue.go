package main

import (
	"fmt"
)

type TaskQueue struct {
	Tasks []Task
}

type Task struct {
}

func (q *TaskQueue) size() int {
	return len(q.Tasks)
}

func (q *TaskQueue) dequeue() (t *Task, err error) {
	var task Task
	if q.size() > 0 {
		task, q.Tasks = q.Tasks[0], q.Tasks[1:]
		return &task, nil
	}
	return nil, fmt.Errorf("Cannot dequeue from an empty queue")
}

func (q *TaskQueue) enqueue(t *Task) {
	q.Tasks = append(q.Tasks, *t)
}
