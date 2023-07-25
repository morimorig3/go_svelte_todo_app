package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := s.srv.Serve(s.l)
		if err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	err := s.srv.Shutdown(context.Background())
	if err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	return eg.Wait()
}
