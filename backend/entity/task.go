package entity

type TaskID int64

type Task struct {
	ID    TaskID `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}

type Tasks []*Task
