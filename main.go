package main

import "log"

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.createAccountTable(); err != nil {
		log.Fatal(err)
	}

	server := NewApiServer(":8888", store)
	server.Run()
}