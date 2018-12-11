package simpleserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Ping is a GET API that responds with a simple text
func Ping(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Pong")
}

// StartHTTPServer starts the server in port 8000 using mux router
func StartHTTPServer() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", Ping)
	address := "0.0.0.0:8000"
	log.Printf("Starting http server on %s\n", address)
	http.ListenAndServe(address, r)
}
