package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"example/db-api/models"

	_ "github.com/lib/pq"
)

// Create a custom SuperGlobal struct which holds a connection pool.
type SuperGlobal struct {
	DbhRw *sql.DB
	DbhRo *sql.DB
}

func main() {
	// Initialise the connection pool.
	//ornamentDbh, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//eggDbh, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Create an instance of SuperGlobal containing the connection pool.
	//eggHolder := &SuperGlobal{DBH: eggDbh}
	//ornamentHolder := &SuperGlobal{DBH: ornamentDbh}

	writeDbh, err := sql.Open("postgres", "postgres://user_rw:pass@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}
	readOnlyDbh, err := sql.Open("postgres", "postgres://user_ro:pass@localhost/bookstore")
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance of SuperGlobal containing the connection pool.
	dbA := &SuperGlobal{DbhRo: readOnlyDbh, DbhRw: writeDbh}

	// Use SuperGlobal.booksIndex as the handler function for the /books route.
	http.HandleFunc("/books", dbA.booksIndex)
	http.ListenAndServe(":3000", nil)
}

// Define booksIndex as a method on SuperGlobal.
func (lDbA *SuperGlobal) booksIndex(w http.ResponseWriter, r *http.Request) {
	// We can now access the connection pool directly in our handlers.
	bks, err := models.AllBooks(lDbA.DbhRo)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
