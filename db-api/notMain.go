package main

import (
	"database/sql"
	"example/db-api/models"
	"fmt"
	"log"
	"net/http"
)

type Env struct {
	db *sql.DB
}

func main() {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{db: db}
	handler := booksIndex(env)

	// Pass the Env struct as a parameter to booksIndex().
	http.Handle("/books", handler)
	http.ListenAndServe(":3000", nil)
}

// Use a closure to make Env available to the handler logic.
func booksIndex(env *Env) http.HandlerFunc {
	//spew.Dump(env)
	return func(w http.ResponseWriter, r *http.Request) {
		bks, err := models.AllBooks(env.db)
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		for _, bk := range bks {
			fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
		}
	}
}
