package pointingPoker

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"pointing-poker/pkg/data"
)

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
		log.Fatal(err)
	}
}

func (s *Server) routes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/pointing-poker", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Group(func(r chi.Router) {
			r.Route("/session", func(r chi.Router) {
				r.Post("/new", s.handleNewSession)
				r.Post("/join/{sessionID}", s.handleJoinSession)
			})
		})
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
