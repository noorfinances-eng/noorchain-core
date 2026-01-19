package health

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	addr string
	srv  *http.Server
	log  *log.Logger
}

func New(addr string) *Server {
	return &Server{
		addr: addr,
		log:  log.New(os.Stdout, "[health] ", log.LstdFlags),
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	s.srv = &http.Server{
		Addr:              s.addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		s.log.Println("listening on", s.addr, "(/healthz)")
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Println("serve error:", err)
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if s.srv == nil {
		return nil
	}
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("health shutdown: %w", err)
	}
	s.log.Println("stopped")
	return nil
}
