package main

import (
	"fmt"
	"learning/fetching-data/fetcher"
)

func main() {
	fmt.Println("We are fetching data")
	facts := fetcher.GetCatFacts()

	for _, fact := range facts {
		fact.PrintText()
	}
}
