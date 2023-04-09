package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/michaelvl/cosign-oci-poc/pkg/version"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
}

func main() {
	log.Printf("version: %s\n", version.Version)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
