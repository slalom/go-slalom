package api

import (
	"github.com/sirupsen/logrus"
	"go-slalom/pkg/version"
	"net/http"
)

func (s *Server) versionHandler(w http.ResponseWriter, r *http.Request) {
	logger.WithFields(logrus.Fields{"api": "version",}).Infof("handling request")

	result := map[string]string{
		"version": version.Version,
		"commit":  version.Revision,
	}
	s.JSONResponse(w, r, result)
}
