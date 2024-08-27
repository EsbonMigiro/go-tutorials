package main

import (
	// "advancedServer/api"
	"advancedServer/api"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var templates *template.Template
var client *redis.Client

func main() {
	var store = sessions.NewCookieStore([]byte("secret-cookie"))
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()

	// srv := api.NewServer(templates)

	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))
	srv := api.NewServer(templates, client, r, store)
	fmt.Println("server listening on port: 4000...")
	http.ListenAndServe(":4000", srv)
}
