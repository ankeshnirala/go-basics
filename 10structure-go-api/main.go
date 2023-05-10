package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ankeshnirala/go/structure-go-api/api"
	"github.com/ankeshnirala/go/structure-go-api/storage"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Println("Server is running on port", *listenAddr)
	log.Fatal(server.Start())

}
