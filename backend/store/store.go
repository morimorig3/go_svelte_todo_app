package store

import (
	"errors"
	"sort"

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

func (ts *TaskStore) Remove(id entity.TaskID) entity.Tasks {
	delete(ts.Tasks, id)
	return ts.All()
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := []*entity.Task{}
	for _, t := range ts.Tasks {
		tasks = append(tasks, &entity.Task{
			ID:     t.ID,
			Title:  t.Title,
		})
	}
	sort.Slice(tasks, func(i, j int) bool {
        return tasks[i].ID < tasks[j].ID
    })
	return tasks
}
