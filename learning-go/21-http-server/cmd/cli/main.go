package main

import (
	"fmt"
	poker "learning_go/21-http-server"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play a game")
	fmt.Println("Type {Name} wins to record a win")

	poker.NewCLI(store, os.Stdin).PlayPoker()
}
