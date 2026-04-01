package pointingPoker

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s *Server) handleNewSession(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	// newSession UUID.
	newSession, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("create new session uuid: %w", err)
	}

	// Store newSession UUID in database.
	err = s.database.Write(ctx, fmt.Sprint(newSession))
	if err != nil {
		return fmt.Errorf("write session %s: %w", newSession, err)
	}

	// Return newSession.
	if _, err = w.Write([]byte(fmt.Sprint(newSession))); err != nil {
		if errors.Is(err, context.Canceled) {
			return fmt.Errorf("write response for new session %s: %w", newSession, err)
		}

		return fmt.Errorf("write response for new session %s: %w", newSession, err)
	}

	return nil
}

func (s *Server) handleJoinSession(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	sessionID := chi.URLParam(r, "sessionID")

	found, err := s.database.Read(ctx, sessionID)
	if err != nil {
		return fmt.Errorf("read session %s: %w", sessionID, err)
	}
	if !found {
		http.Error(w, "sessionID not found", http.StatusNotFound)
		return nil
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
