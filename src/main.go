package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, r *http.Request) {
		file, err := os.Open("public" + r.URL.Path)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}

		defer file.Close()

		var ContentType string

		switch  {
		case strings.HasSuffix(r.URL.Path, "css"):
			ContentType = "text/css"
		case strings.HasSuffix(r.URL.Path, "html"):
			ContentType = "text/html"
		case strings.HasSuffix(r.URL.Path, "png"):
			ContentType = "text/png"
		default:
			ContentType = "text/plain"
		}

		response.Header().Add("Content-Type", ContentType)
		io.Copy(response, file)
	})

	http.ListenAndServe(":8000", nil)
}
