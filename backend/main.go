package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/morimorig3/go_svelte_todo_app/backend/config"
)

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	mux := NewMux()
	s := NewServer(l, mux)
	return s.Run(ctx)
}

func main() {
	err := run(context.Background())
	if err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}
