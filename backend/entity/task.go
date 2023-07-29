package entity

type TaskID int64

type Task struct {
	ID      TaskID     `json:"id"`
	Title   string     `json:"title"`
}

type Tasks []*Task
