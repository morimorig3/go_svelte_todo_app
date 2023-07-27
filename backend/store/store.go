package store

import (
	"errors"
	"fmt"

	"github.com/morimorig3/go_svelte_todo_app/backend/entity"
)

var (
	Tasks = &TaskStore{
		Tasks: map[entity.TaskID]*entity.Task{},
	}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	LastID entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) Remove(id entity.TaskID) {
	fmt.Printf("%v", ts.Tasks)
	delete(ts.Tasks, id)
	fmt.Printf("%v", ts.Tasks)
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := []*entity.Task{}
	for _, t := range ts.Tasks {
		tasks = append(tasks, &entity.Task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	return tasks
}
