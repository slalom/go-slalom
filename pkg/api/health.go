package api

import (
	"net/http"
	"sync/atomic"
)

var (
	healthy int32
	ready   int32
)

func (s *Server) enableProbes() {
	atomic.StoreInt32(&healthy, 1)
	atomic.StoreInt32(&ready, 1)
}

func (s *Server) disableProbes() {
	atomic.StoreInt32(&healthy, 0)
	atomic.StoreInt32(&ready, 0)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		s.JSONResponse(w, r, map[string]string{"status": "OK"})
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

func (s *Server) readyHandler(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&ready) == 1 {
		s.JSONResponse(w, r, map[string]string{"status": "OK"})
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

func (s *Server) enableReadyHandler(w http.ResponseWriter, r *http.Request) {
	atomic.StoreInt32(&ready, 1)
	w.WriteHeader(http.StatusAccepted)
}

func (s *Server) disableReadyHandler(w http.ResponseWriter, r *http.Request) {
	logger.Warn("disabling ready probe")
	atomic.StoreInt32(&ready, 0)
	w.WriteHeader(http.StatusAccepted)
}

func (s *Server) disableHealthHandler(w http.ResponseWriter, r *http.Request) {
	logger.Warn("disabling health probe")
	atomic.StoreInt32(&healthy, 0)
	w.WriteHeader(http.StatusAccepted)
}
