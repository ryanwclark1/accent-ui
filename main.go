package main

import (
	"accent-ui/handler"
	"log"
	"log/slog"
	"net/http"
	"embed"
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
	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))

	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("application running", "port", port )
	log.Fatal(http.ListenAndServe(port, router))
}

func initEverything() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}