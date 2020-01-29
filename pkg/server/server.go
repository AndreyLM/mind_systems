package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/andreylm/mind_systems/pkg/db"
	"github.com/andreylm/mind_systems/pkg/ent"
	customHanlders "github.com/andreylm/mind_systems/pkg/server/handlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server - server
type Server struct {
	host   string
	port   string
	router *mux.Router
	db     *ent.Client
}

// NewServer - creates new server
func NewServer(host, port string) *Server {

	s := &Server{
		host:   host,
		port:   port,
		router: mux.NewRouter(),
	}

	return s
}

// Start - startign server
func (s *Server) Start() error {
	if err := s.init(); err != nil {
		return err
	}
	defer s.db.Close()

	srv := &http.Server{
		Handler:      s.getHandler(),
		Addr:         fmt.Sprintf("%s:%s", s.host, s.port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

func (s *Server) init() (err error) {
	if s.db, err = db.CreateConnection(); err != nil {
		return err
	}

	s.router.HandleFunc("/create-url", customHanlders.Shortener(fmt.Sprintf("%s:%s", s.host, s.port), s.db))
	s.router.HandleFunc("/{urlToken}", customHanlders.URLResolver(s.db))
	return nil
}

func (s *Server) getHandler() http.Handler {
	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control"}),
		handlers.ExposedHeaders([]string{"*"}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(s.router))

	return handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)
}
