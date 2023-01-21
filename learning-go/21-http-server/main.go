package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore(), sync.Mutex{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
