package main

import (
	"fmt"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "¡Bienvenido a la ruta principal! 🚀")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status": "OK", "env": "github-actions-test"}`)
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Esta es una ruta de prueba para validar flujos de CI/CD.")
	})

	mux.HandleFunc("/beda", func(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
		w.WriteHeader(http.StatusBad)
		fmt.Fprintln(w, "Tengo 24 anos de edad")
=======
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Esta soy yo")
>>>>>>> b4ac30661cd226e828dda517898fd350c03b0bec
	})

	return mux
}