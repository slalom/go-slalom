package api

import (
	"net/http"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/tredfield/go-slalom/pkg/version"
)

func (s *Server) infoHandler(w http.ResponseWriter, r *http.Request) {
	logger.WithFields(logrus.Fields{"api": "info"}).Infof("handling request")

	data := struct {
		Hostname     string `json:"hostname"`
		Version      string `json:"version"`
		Revision     string `json:"revision"`
		Color        string `json:"color"`
		Message      string `json:"message"`
		GOOS         string `json:"goos"`
		GOARCH       string `json:"goarch"`
		Runtime      string `json:"runtime"`
		NumGoroutine string `json:"num_goroutine"`
		NumCPU       string `json:"num_cpu"`
	}{
		Hostname: "TODO",
		Version:  version.Version,
		Revision: version.Revision,
		//Color:        s.config.UIColor,
		//Message:      s.config.UIMessage,
		GOOS:         runtime.GOOS,
		GOARCH:       runtime.GOARCH,
		Runtime:      runtime.Version(),
		NumGoroutine: strconv.FormatInt(int64(runtime.NumGoroutine()), 10),
		NumCPU:       strconv.FormatInt(int64(runtime.NumCPU()), 10),
	}

	s.JSONResponse(w, r, data)
}
