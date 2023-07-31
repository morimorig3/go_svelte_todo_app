package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/morimorig3/go_svelte_todo_app/backend/config"
	"github.com/morimorig3/go_svelte_todo_app/backend/handler"
	"github.com/morimorig3/go_svelte_todo_app/backend/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	r := &store.Repository{}
	at := &handler.AddTask{
		DB:        db,
		Repo:      r,
		Validator: v,
	}
	mux.Post("/task/add", at.ServeHTTP)
	lt := &handler.ListTask{
		DB:   db,
		Repo: r,
	}
	mux.Get("/task/list", lt.ServeHTTP)
	rt := &handler.RemoveTask{
		DB:        db,
		Repo:      r,
		Validator: v,
	}
	mux.Post("/task/remove", rt.ServeHTTP)
	return mux, cleanup, nil
}
