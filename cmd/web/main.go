package main

import (
	"log"
	"net/http"
)

// WEB app:
// 1. servemux(router)
// 2. handler(funcs)
// 3. web server

// the web server receives http request -> pass the request on the servemux
//  -> the servemux checks the URL -> sends the request to the matching handler

// CURL:

// curl -i -X POST <URL>
// curl -i -X PUT <URL>

func main() {

	// 1. new servemux
	mux := http.NewServeMux()

	// http.HandleFunc without "servemux" ==  DefaultServeMux, global var
	// -> bad for security

	// a) subtree path: "/" at the end
	// matches if the start of URL matches
	// that's why there's might be a CATCH-ALL
	mux.HandleFunc("/", home)

	// b) fixed path: NO "/" at the end
	// only matches with the exact URL (NOT: /snippet/view/)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// 3. web server
	log.Println("Starting server on: 4000")
	err := http.ListenAndServe("localhost:4000", mux)
	// a) ":4000" ALL network interfaces
	// b) "localhost:4000" NOT ALL
	if err != nil {
		log.Fatal(err)
	}
}
