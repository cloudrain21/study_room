package main

import (
	"github.com/cloudrain21/poker"
	"log"
	"net/http"
)

func main() {
	database, closeFunc, err := poker.NewFileSystemPlayerStoreFromFile("testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	server := poker.NewPlayerServer(database)

	err = http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("can't start server : %s", err)
	}
}
