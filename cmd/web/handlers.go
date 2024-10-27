package main

import (
	"html/template"
	"log"
	"net/http"
)

// handler #1: homepage

func home(w http.ResponseWriter, r *http.Request) {

	// avoiding CATCH-ALL:
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("/Users/daryakuznetsova/Desktop/WEB/projects/snippet/ui/html/pages/home.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "ParseInternalServerError", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "ExecuteInternalServerError", http.StatusInternalServerError)
	}
	w.Write([]byte("Hello from Snippet ^.^"))
}
