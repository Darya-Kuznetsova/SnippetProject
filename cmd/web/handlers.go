package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

// handler #2: display a specific snippet

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

// handler #3: creating a new snippet

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	// POST method only allowed:

	if r.Method != http.MethodPost { // "POST"

		// Header().Set():
		// before WriteHeader() + Write() (or http.Error())

		w.Header().Set("Allow", http.MethodPost)

		// 1. WriteHeader():
		// a) only possible to call ONCE
		// b) has to stay BEFORE Write()

		// w.WriteHeader(http.StatusMethodNotAllowed) // 405
		// w.Write([]byte("Method Not Allowed"))

		// OR:

		// 2. http.Error():
		// a) most common ^^

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // 405
		return
	}
	w.Write([]byte("Create a new snippet ^^"))

}
