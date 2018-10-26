package basic

import (
	"fmt"
	"log"
	"net/http"
)

func Goserver() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}