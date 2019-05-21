package api

import (
	"net/http"
	"text/template"
	"time"
)

func (s *Server) welcomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/welcome.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("templates/welcome.html" + err.Error()))
		return
	}

	data := struct {
		Name string
		Time string
	}{
		Name: "slalom",
		Time: time.Now().Format(time.Stamp),
	}

	if name := r.FormValue("name"); name != "" {
		data.Name = name
	}

	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, "templates/welcome.html"+err.Error(), http.StatusInternalServerError)
	}
}
