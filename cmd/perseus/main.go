package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"tildegit.org/andinus/perseus/handler/web"
	"tildegit.org/andinus/perseus/storage"
)

func main() {
	db := storage.Init()
	defer db.Conn.Close()

	envPort := os.Getenv("PERSEUS_PORT")
	if envPort == "" {
		envPort = "8080"
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf("127.0.0.1:%s", envPort),
		WriteTimeout: 8 * time.Second,
		ReadTimeout:  8 * time.Second,
	}

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		web.RegisterHandler(w, r, db)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		web.LoginHandler(w, r, db)
	})

	log.Printf("perseus: listening on port %s...", envPort)
	log.Fatal(srv.ListenAndServe())
}
