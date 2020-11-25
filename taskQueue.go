package main

type TaskQueue struct {
	Tasks []Task
}

type Task struct {
}

func (q *TaskQueue) size() (int) {
	return len(q.Tasks)
}
