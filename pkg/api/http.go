package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// JSONResponse takes a provided result and marshals to json applying appropriate headers
func (s *Server) JSONResponse(w http.ResponseWriter, r *http.Request, result interface{}) {
	body, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error("JSON marshal failed", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	bytes, err := prettyJSON(body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error("JSON prettyJSON failed", err)
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error("Failed to write http response", err)
	}
}

func prettyJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
