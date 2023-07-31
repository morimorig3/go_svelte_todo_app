package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/morimorig3/go_svelte_todo_app/backend/entity"
	"github.com/morimorig3/go_svelte_todo_app/backend/store"
)

type RemoveTask struct {
	DB        *sqlx.DB
	Repo      *store.Repository
	Validator *validator.Validate
}

func (rt *RemoveTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		ID entity.TaskID `json:"id" validate:"required"`
	}
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		RespondJSON(ctx, w, &ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	err = rt.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	rt.Repo.RemoveTask(ctx, rt.DB, &b.ID)
	RespondJSON(ctx, w, nil, http.StatusOK)
}
