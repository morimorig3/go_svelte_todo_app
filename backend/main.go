package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/morimorig3/go_svelte_todo_app/backend/config"
	"golang.org/x/sync/errgroup"
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

	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := s.Serve(l)
		if err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})
	<-ctx.Done()
	err = s.Shutdown(context.Background())
	if err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	return eg.Wait()
}

func main() {
	err := run(context.Background())
	if err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}
