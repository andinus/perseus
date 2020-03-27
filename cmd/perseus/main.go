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

	envPort, exists := os.LookupEnv("PERSEUS_PORT")
	if !exists {
		envPort = "8080"
	}
	addr := fmt.Sprintf("127.0.0.1:%s", envPort)

	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: 8 * time.Second,
		ReadTimeout:  8 * time.Second,
	}

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		web.HandleRegister(w, r, db)
	})

	log.Printf("main/main.go: listening on port %s...", envPort)
	log.Fatal(srv.ListenAndServe())
}
