package handler

import (
	"net/http"

	"github.com/morimorig3/go_svelte_todo_app/backend/entity"
	"github.com/morimorig3/go_svelte_todo_app/backend/store"
)

type ListTask struct {
	Store *store.TaskStore
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks := lt.Store.All()
	rsp := []entity.Task{}
	for _, t := range tasks {
		rsp = append(rsp, entity.Task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
