package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(v)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]string{"mensagem": "Olá da API Go na Nextcluster!", "go": runtime.Version(), "porta": port})
	})
	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]string{"status": "ok"})
	})
	r.Get("/ola/{nome}", func(w http.ResponseWriter, req *http.Request) {
		writeJSON(w, map[string]string{"mensagem": "Olá, " + chi.URLParam(req, "nome") + "!"})
	})

	log.Printf("ouvindo na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
