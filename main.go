package main

import (
	"accent-ui/handler"
	"accent-ui/pkg/accentapi"

	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(handler.WithUser)

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.Make(handler.HandleHomeIndex))
	router.Get("/login", handler.Make(handler.HandleLoginIndex))
	router.Post("/login", handler.Make(handler.HandleLoginCreate))

	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("application running", "port", port )
	log.Fatal(http.ListenAndServe(port, router))
}

// initEverything initializes everything needed for the application to run
// This includes loading environment variables and initializing the accent API client
// This might not be best practice, could intialize accent API client in another manner
func initEverything() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return accentapi.Init()
	// return accentapi.Init()
}