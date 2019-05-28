package api

import (
	"net/http"
	"text/template"
)

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("templates/index.html" + err.Error()))
		return
	}
	if err := tmpl.ExecuteTemplate(w, "layout", ""); err != nil {
		logger.Errorf("error executing templates: %s", err.Error())
		http.Error(w, "templates/welcome.html"+err.Error(), http.StatusInternalServerError)
	}
}
