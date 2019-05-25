package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, r *http.Request) {
		http.ServeFile(response, r, "public"+r.URL.Path)
	})

	http.ListenAndServe(":8000", nil)
}
