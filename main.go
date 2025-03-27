package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"htmx-go-workshop/internal/pokedex"
)

type Templates struct {
	templates *template.Template
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Templates) render(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := t.templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	tmpl := newTemplate()
	mux := http.NewServeMux()

	// GET endpoint to get a list of Pokemon
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.render(w, "index.html", pokedex.GetPokedex())
	})

	// GET endpoint to list all Pokemon that match the query
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		data := struct {
			PokemonEntries []pokedex.PokemonEntry
		}{
			PokemonEntries: pokedex.SearchPokedex(query),
		}
		tmpl.render(w, "entries.html", data)
	})

	// GET endpoint to get a specific pokemon based on their pokedex number
	mux.HandleFunc("/pokemon/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/pokemon/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid Pokemon ID", http.StatusBadRequest)
			return
		}
		entry, err := pokedex.GetPokemonEntry(id)
		if err != nil {
			http.Error(w, "Pokemon not found", http.StatusNotFound)
			return
		}
		tmpl.render(w, "entry.html", entry)
	})

	addr := ":3000"
	log.Printf("Starting server at Port %s", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
