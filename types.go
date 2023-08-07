package main

import "math/rand"

type Account struct {
	Id         int `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	CardNumber int64 `json:"number"`
	Balance    int64 `json:"balance"`
}

func NewAccout(firstName, lastName string) *Account {
	return &Account{
		Id: rand.Intn(10000),
		FirstName: firstName,
		LastName: lastName,
		CardNumber: int64(rand.Intn(1000000)),
	}
}