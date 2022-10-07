// server1 is a minimum echo server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request of route / will call the func handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
