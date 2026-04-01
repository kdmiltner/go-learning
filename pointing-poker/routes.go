package pointingPoker

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"pointing-poker/pkg/data"
)

type appHandler func(http.ResponseWriter, *http.Request) error
type Server struct {
	database data.ReadWriter
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	s.bootstrapServer()
	s.routes()
}

func (s *Server) bootstrapServer() {
	var err error
	s.database, err = data.NewDatabase(data.DatabaseCSV)
	if err != nil {
		panic(err)
	}
}

func (s *Server) routes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/pointing-poker", func(r chi.Router) {
		r.Use(middleware.Recoverer)
		r.Group(func(r chi.Router) {
			r.Route("/session", func(r chi.Router) {
				r.Post("/new", s.withErrorHandling(s.handleNewSession))
				r.Post("/join/{sessionID}", s.withErrorHandling(s.handleJoinSession))
			})
		})
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World"))
	})
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func (s *Server) withErrorHandling(next appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err == nil {
			return // Return early if no error occurred.
		}

		switch {
		case errors.Is(err, context.Canceled):
			// Client disconnected/canceled; usually don't try to write a response.
			log.Printf("request canceled: %v", err)
			return
		case errors.Is(err, context.DeadlineExceeded):
			http.Error(w, "request timed out", http.StatusGatewayTimeout)
			return
		default:
			log.Printf("request failed: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}
}
