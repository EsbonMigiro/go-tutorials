package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>This is my Awesome website</h1>")
}

func main() {

	http.HandleFunc("/", handlerFunc)
	fmt.Println("server running on port 3000")
	http.ListenAndServe(":4000", nil)

}
