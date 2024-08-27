package main

import "net/http"

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(" introduction"))
	})

	http.ListenAndServe(":8080", nil)
}
