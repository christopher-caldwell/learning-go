package main

import (
	"log"
	"net/http"

	"oapigen/account"
	"oapigen/pets"
	petstore "oapigen/petstore"
)

type handler struct {
	pets.PetsService
	account.AccountService
}

var _ petstore.Handler = (*handler)(nil)

func main() {
	service := &handler{
		PetsService:    pets.Service,
		AccountService: account.Service,
	}
	srv, err := petstore.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
