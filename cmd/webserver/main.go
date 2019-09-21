package main

import (
	"log"
	"net/http"

	poker "github.com/odraude1022/http-server"
)

const dbFileName = "game.db.json"

func main() {

	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server, _ := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
