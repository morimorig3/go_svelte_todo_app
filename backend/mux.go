package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/morimorig3/go_svelte_todo_app/backend/handler"
	"github.com/morimorig3/go_svelte_todo_app/backend/store"
)

func NewMux() http.Handler {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	v := validator.New()
	at := &handler.AddTask{
		Store:     store.Tasks,
		Validator: v,
	}
	mux.Post("/task/add", at.ServeHTTP)
	lt := &handler.ListTask{
		Store: store.Tasks,
	}
	mux.Get("/task/list", lt.ServeHTTP)
	rt := &handler.RemoveTask{
		Store: store.Tasks,
		Validator: v,
	}
	mux.Post("/task/remove", rt.ServeHTTP)
	return mux
}
