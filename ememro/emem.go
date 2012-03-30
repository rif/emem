package hello

import (    
    "net/http"
    "html/template"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	signTemplate, err := template.ParseFiles("index.html")
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    err = signTemplate.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
