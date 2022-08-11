package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error terjadi", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini adalah GET"))
	case "POST":
		w.Write([]byte("Ini adalah POST"))
	default:
		http.Error(w, "Error terjadi", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error terjadi", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error terjadi", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error terjadi", http.StatusBadRequest)
}

func Prosses(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error terjadi", http.StatusInternalServerError)
			return
		}

		task := r.Form.Get("task")
		assignee := r.Form.Get("assignee")
		deadline := r.Form.Get("deadline")

		data := map[string]interface{}{
			"task":     task,
			"assignee": assignee,
			"deadline": deadline,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error terjadi", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error terjadi", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error terjadi", http.StatusBadRequest)

}
