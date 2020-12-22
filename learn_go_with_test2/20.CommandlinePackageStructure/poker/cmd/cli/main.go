package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudrain21/poker"
)

func main() {
	fmt.Println("Let's play poker")

	store, closeFunc, err := poker.NewFileSystemPlayerStoreFromFile("testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	poker.NewCLI(store,os.Stdin).PlayPoker()
}
