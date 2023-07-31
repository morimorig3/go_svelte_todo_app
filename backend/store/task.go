package store

import (
	"context"

	"github.com/morimorig3/go_svelte_todo_app/backend/entity"
)

func (r *Repository) ListTasks(ctx context.Context, db Queryer) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT id, title FROM task;`
	err := db.SelectContext(ctx, &tasks, sql)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) AddTask(ctx context.Context, db Execer, t *entity.Task) error {
	sql := `INSERT INTO task(title) VALUES(?)`

	result, err := db.ExecContext(ctx, sql, t.Title)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}

func (r *Repository) RemoveTask(ctx context.Context, db Execer, id *entity.TaskID) error {
	sql := `DELETE FROM task WHERE id = ?`

	_, err := db.ExecContext(ctx, sql, id)
	if err != nil {
		return err
	}

	return nil
}
