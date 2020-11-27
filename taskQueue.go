package main

import (
	"fmt"
	"time"
	"math/rand"
)

type TaskQueue struct {
	Tasks []Task
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

func (q *TaskQueue) end() (t *Task, err error) {
	var task Task
	if q.size() > 0 {
		task = q.Tasks[len(q.Tasks)-1]
		return &task, nil
	}
	return nil, fmt.Errorf("Cannot dequeue from an empty queue")
}

type TaskData struct {
	MaxRuntime int32
	TimeRemaining int32
}

type TaskStatus int

const (
	RUNNING TaskStatus = iota
	FAILED
	COMPLETED
	TIMEDOUT
)

type Task struct {
	ID          string
	Status      TaskStatus
	IsCompleted bool
	SubmitTime  time.Time
	Data    TaskData
}

func (t *Task) updateRemainingTime () {
	if t.IsCompleted {
		return
	}
	r := rand.New(rand.NewSource(99))
	value := r.Int31n(5)
	t.Data.TimeRemaining -= value
	if t.Data.TimeRemaining <= 0 {
		t.IsCompleted = true
	}
}

func NewTask(name string, runtime int32) (t *Task) {
	return &Task{
		ID: name,
		Status: RUNNING,
		IsCompleted: false,
		SubmitTime: time.Now(),
		Data: TaskData{
			MaxRuntime: runtime,
			TimeRemaining: runtime,
		},
	}
}
