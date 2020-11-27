package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

type TaskQueue struct {
	sync.Mutex
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

func main() {

	q := TaskQueue{}

	r := rand.New(rand.NewSource(123))
	q.enqueue(NewTask("task1", r.Int31n(100000)))
	q.enqueue(NewTask("task2", r.Int31n(100000)))
	q.enqueue(NewTask("task3", r.Int31n(100000)))
	q.enqueue(NewTask("task4", r.Int31n(100000)))
	q.enqueue(NewTask("task5", r.Int31n(100000)))
	q.enqueue(NewTask("task6", r.Int31n(100000)))
	q.enqueue(NewTask("task7", r.Int31n(100000)))
	q.enqueue(NewTask("task8", r.Int31n(100000)))
	q.enqueue(NewTask("task9", r.Int31n(100000)))
	q.enqueue(NewTask("task0", r.Int31n(100000)))

	shd := make(chan int, 1) // schedular done
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go Cleaner(&q, shd, wg)
	for {
		q.Lock()
		task, err := q.dequeue()
		if err != nil {
			break
		}
		task.updateRemainingTime()
		q.enqueue(task)
		q.Unlock()
	}
	shd <- 1
	fmt.Println("Exiting")
	wg.Wait()
}

func Cleaner(q *TaskQueue, quit chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-quit:
			wg.Done()
			return
		default:
			fmt.Println("Cleaner Invoked")
			q.Lock()
			last_task, err := q.end()
			if err != nil {
				break
			}
			for {
				task, err := q.dequeue()
				if err != nil {
					break
				}
				if ! task.IsCompleted {
					q.enqueue(task)
				} else {
					fmt.Println(task.ID, "task cleared")
				}
				if task.ID == last_task.ID {
					break
				}
			}
			q.Unlock()
			time.Sleep(100*time.Millisecond)
		}
	}
}
