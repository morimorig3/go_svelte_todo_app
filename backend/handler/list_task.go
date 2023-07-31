package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/morimorig3/go_svelte_todo_app/backend/entity"
	"github.com/morimorig3/go_svelte_todo_app/backend/store"
)

type ListTask struct {
	DB   *sqlx.DB
	Repo *store.Repository
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := lt.Repo.ListTasks(ctx, lt.DB)
	if err != nil {
		RespondJSON(ctx, w, &ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []entity.Task{}
	for _, t := range tasks {
		rsp = append(rsp, entity.Task{
			ID:    t.ID,
			Title: t.Title,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
