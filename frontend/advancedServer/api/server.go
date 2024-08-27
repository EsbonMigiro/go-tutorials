package api

import (
	"fmt"
	htmlTemplate "html/template"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	*mux.Router

	templates *htmlTemplate.Template
	client    *redis.Client
	store     *sessions.CookieStore
}

func NewServer(templates *htmlTemplate.Template, client *redis.Client, router *mux.Router, store *sessions.CookieStore) *Server {
	s := &Server{
		Router:    router,
		templates: templates,
		client:    client,
		store:     store,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/hello-world", s.helloWordUrl()).Methods("GET")
	s.HandleFunc("/index", s.indexHandler()).Methods("GET")
	s.HandleFunc("/test", s.testGetHandler()).Methods("GET")
	s.HandleFunc("/index", s.AuthRequiredMiddleWare(s.indexPostHandler())).Methods("POST")
	s.HandleFunc("/login", s.loginGetHandler()).Methods("GET")
	s.HandleFunc("/login", s.loginPostHandler()).Methods("POST")
	s.HandleFunc("/register", s.registerGetHandler()).Methods("GET")
	s.HandleFunc("/register", s.registerPostHandler()).Methods("POST")
	s.HandleFunc("/", s.AuthRequiredMiddleWare(s.landingGetHandler())).Methods("GET")

}

func (s *Server) AuthRequiredMiddleWare(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.store.Get(r, "sessions")
		_, ok := session.Values["username"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func (s *Server) helloWordUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello World</h1>")
	}
}

func (s *Server) landingGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// session, _ := s.store.Get(r, "session")
		// _, ok := session.Values["username"]

		// if !ok {
		// 	http.Redirect(w, r, "/login", 302)

		// }
		fmt.Fprint(w, "<h1>landed</h1>")
	}
}

func (s *Server) indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		comment, err := s.client.LRange("comments", 0, 10).Result()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal Server Error"))
			return
		}
		s.templates.ExecuteTemplate(w, "index.html", comment)
	}
}

func (s *Server) indexPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		comment := r.PostForm.Get("comment")
		err := s.client.LPush("comments", comment).Err()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal Server Error"))
		}
		http.Redirect(w, r, "/index", 302)
	}
}

func (s *Server) loginGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.templates.ExecuteTemplate(w, "login.html", nil)
	}
}

func (s *Server) loginPostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")

		hash, err := s.client.Get("user:" + username).Bytes()
		fmt.Println("hash", hash)
		if err == redis.Nil {
			s.templates.ExecuteTemplate(w, "login.html", "Uknown user")
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal Server Error hash"))
			return
		}

		errr := bcrypt.CompareHashAndPassword(hash, []byte(password))
		if errr != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			// w.Write([]byte("internal Server Error CompareHashAndPassword"))

			s.templates.ExecuteTemplate(w, "login.html", "Invalid Login")
			return
		}

		session, _ := s.store.Get(r, "session")
		session.Values["username"] = username
		session.Save(r, w)

		http.Redirect(w, r, "/index", 302)

	}
}

func (s *Server) testGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.store.Get(r, "session")
		untyped, ok := session.Values["username"]

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}
		username, ok := untyped.(string)

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		w.Write([]byte(username))

	}
}

func (s *Server) registerGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.templates.ExecuteTemplate(w, "register.html", nil)
	}

}
func (s *Server) registerPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")

		cost := bcrypt.DefaultCost

		hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		errr := s.client.Set("user:"+username, hash, 0).Err()
		if errr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal Server Error"))
			return

		}

		http.Redirect(w, r, "/login", 302)
	}
}

// import (
// 	"html/template"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/go-redis/redis"
// )

// type Server struct {
// 	*mux.Router

// 	templates *template.Template
// }
// var client *redis.Client

// func NewServer(templates *template.Template) *Server {
// 	s := &Server{
// 		Router:    mux.NewRouter(),
// 		templates: templates,
// 	}
// 	s.routes()
// 	return s
// }

// func (s *Server) routes() {
// 	s.HandleFunc("/hello", s.createHelloWorldRoute()).Methods("GET")
// }

// func (s *Server) createHelloWorldRoute() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		s.templates.ExecuteTemplate(w, "hello.html", nil)
// 	}

// }
