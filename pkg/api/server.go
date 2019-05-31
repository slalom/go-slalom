package api

import (
	"context"
	"net/http"

	"github.com/slalom/go-slalom/pkg/signals"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

var (
	logger = logrus.WithFields(logrus.Fields{"service": "api"})
)

const (
	staticDir   = "/static/"
	templateDir = "/templates/"
	port        = "8080"
)

// Server provides go-slalom server
type Server struct {
	router *mux.Router
}

// NewServer constructs new Server
func NewServer() *Server {
	srv := &Server{
		router: mux.NewRouter(),
	}

	return srv
}

func (s *Server) registerHandlers() {
	s.router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	s.router.PathPrefix(templateDir).Handler(http.StripPrefix(templateDir, http.FileServer(http.Dir("."+templateDir))))
	s.router.HandleFunc("/", s.indexHandler).HeadersRegexp("User-Agent", "^Mozilla.*").Methods("GET")
	s.router.HandleFunc("/", s.infoHandler).Methods("GET")
	s.router.HandleFunc("/welcome", s.welcomeHandler).Methods("GET")
	s.router.HandleFunc("/version", s.versionHandler).Methods("GET")
	s.router.HandleFunc("/health", s.healthHandler).Methods("GET")
	s.router.HandleFunc("/health/disable", s.disableHealthHandler).Methods("POST")
	s.router.HandleFunc("/ready", s.readyHandler).Methods("GET")
	s.router.HandleFunc("/ready/enable", s.enableReadyHandler).Methods("POST")
	s.router.HandleFunc("/ready/disable", s.disableReadyHandler).Methods("POST")
}

// ListenAndServe starts server and begins taking requests
func (s *Server) ListenAndServe() {
	logger.Infof("Starting server...")

	// setup channel to handle interrupt signals
	stopCh := signals.SignalHandler()

	// create routes
	s.registerHandlers()

	// create http server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: s.router,
	}

	// run in background
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// Error starting or closing listener:
			logger.Errorf("HTTP server ListenAndServe: %v", err)
		}
	}()

	// signal Kubernetes the server is ready to receive traffic
	s.enableProbes()

	// wait for interrupt signal
	<-stopCh

	logger.Warn("stopping server")

	// fail health and readiness probes now
	s.disableProbes()

	// graceful shutdown
	if err := srv.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		logger.Errorf("HTTP server Shutdown: %v", err)
	}
}
