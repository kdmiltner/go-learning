package pointingPoker

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (s *Server) handleNewSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println(ctx)

	// newSession UUID.
	newSession, err := uuid.NewUUID()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error creating new session: %v", err)))
		return
	}

	// Store newSession UUID in database.
	err = s.database.Write(fmt.Sprint(newSession))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error: %v", err)))
		return
	}

	// Return newSession.
	w.Write([]byte(fmt.Sprint(newSession)))
}

func (s *Server) handleJoinSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println(ctx)
	sessionID := chi.URLParam(r, "sessionID")

	found, err := s.database.Read(sessionID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error finding sessionID: %v", err)))
		return
	}
	if !found {
		w.WriteHeader(404)
		w.Write([]byte("sessionID not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
