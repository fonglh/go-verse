package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Println(r.URL.RawQuery)
}

func main() {
	http.HandleFunc("/passage", handler)
	http.ListenAndServe(":8080", nil)
}
