package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Ping(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Pong")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", Ping)
	address := "0.0.0.0:80"
	// caddyInst := webserver.StartServer()
	fmt.Printf("Starting http server on %s\n", address)
	http.ListenAndServe(address, r)
	// caddyInst.Wait()
}
