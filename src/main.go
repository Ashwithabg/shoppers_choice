package main

import (
	"github.com/shoppers_choice/src/viewmodel"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	templates := populateTemplates()


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		template := templates[requestedFile+".html"]
		var context interface{}

		switch requestedFile {
		case "shop":
			context = viewmodel.NewShop()
		default:
			context = viewmodel.NewBase()
		}

		if template != nil {
			err := template.Execute(w, context)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	for _, fileInfo := range fileInfos {
		file, err := os.Open(basePath + "/content/" + fileInfo.Name())
		if err != nil {
			panic("Failed to open template '" + fileInfo.Name() + "'")
		}

		content, err := ioutil.ReadAll(file)
		if err != nil {
			panic("Failed to read content from file '" + fileInfo.Name() + "'")
		}

		file.Close()

		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fileInfo.Name() + "' as template")
		}

		result[fileInfo.Name()] = tmpl
	}
	return result
}
