package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	err := OpenDatabase()

	if err != nil {
		log.Printf("error while connecting to the database: %v", err)
	}
	// defer CloseDatabase()

	server()

	fmt.Println("server listening on port: 4000...")

	http.ListenAndServe(":4000", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server working"))
}

func server() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/create-book", createBook)
	http.HandleFunc("/get", getAll)
}

type CreateBookBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var body CreateBookBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Badrequest in decoding object %v", err)
		return
	}
	log.Printf(" body.Name: %v", body.Name)

	insertBook := func() error {
		query := `
		INSERT INTO books (name, description, author) VALUES ($1, $2, $3)
		`
		return DB.QueryRow(query, body.Name, body.Description, body.Author).Err()

	}
	if err := insertBook(); err != nil {
		if strings.Contains(err.Error(), `relation "books" does not exist`) {
			if err := createTable(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("error while in createTable after checking 1 %v", err)
				return
			}
			if err := insertBook(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("error occured while inserting the book %v", err)
				return
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("error while in createTable after checking 2 %v", err)
			return
		}

	}

	//==========================================================================
	w.WriteHeader(http.StatusCreated)
}

func getAll(w http.ResponseWriter, _ *http.Request) {
	var books []Book

	rows, errr := DB.Query("SELECT id, name, description, author FROM books;")
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error during rows %v", errr)

		return
	}

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Name, &book.Description, &book.Author); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("error during Scan %v", err)
			return
		}
		books = append(books, book)
	}
	j, err := json.Marshal(books)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error during getAll %v", err)
		return
	}
	w.Write(j)
}

func createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS books(
		id SERIAL PRIMARY KEY,
		name VARCHAR(100),
		description TEXT,
		author VARCHAR(100)
	);
	`
	_, err := DB.Exec(query)

	if err != nil {
		return err
	}
	return nil
}
