package main

import (
	"tildegit.org/andinus/perseus/storage"
)

func main() {
	db := storage.Init()
	defer db.Conn.Close()
}
