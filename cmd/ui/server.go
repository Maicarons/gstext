package ui

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/Maicarons/gstext"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

//go:embed static
var staticFiles embed.FS

func Start() {
	r := chi.NewRouter()

	// Serve static files
	staticFS, _ := fs.Sub(staticFiles, "static")
	staticFileServer := http.FileServer(http.FS(staticFS))
	r.Handle("/*", http.StripPrefix("/", staticFileServer))

	// Serve index.html
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(staticFS, "index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln(err.Error())
		}
		OrgFile, err := os.ReadFile("j.json")
		Obj, err := gstext.JSONUnmarshalGameText(OrgFile)
		if err != nil {
			log.Fatalln(err.Error())
		}
		err = tmpl.Execute(w, Obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln(err.Error())
		}
	})

	// Save game text
	r.Post("/save", func(w http.ResponseWriter, r *http.Request) {
		var gameText gstext.GameText
		fmt.Println(r.Body)
		err := json.NewDecoder(r.Body).Decode(&gameText)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalln(err.Error())
		}
		marshal, err := gameText.JSONMarshal()
		if err != nil {
			log.Fatalln(err.Error())
		}
		err = os.WriteFile("j.json", marshal, 0777)
		// Process and save the game text as needed

		fmt.Fprint(w, "Game text saved successfully!")
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
