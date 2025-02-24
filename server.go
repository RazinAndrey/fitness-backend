package fitnessbackend

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// старт
func (s *Server) Run(port string,  handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler: handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe() // бесконечный цикл for => слушает все входящие запросы для последующей обработки
}

// стоп
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
