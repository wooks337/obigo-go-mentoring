package main

import (
	"fmt"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

var sessionKey = "session_test_key"

func main() {
	n := negroni.Classic()
	store := cookiestore.New([]byte(sessionKey))
	n.Use(sessions.Sessions(sessionKey, store))

	mux := http.NewServeMux()
	mux.HandleFunc("/create", create)
	mux.HandleFunc("/read", read)
	mux.HandleFunc("/delete", delete)
	mux.HandleFunc("/check", check)

	n.UseHandler(mux)
	n.Run(":3000")
}

func create(w http.ResponseWriter, req *http.Request) {
	s := sessions.GetSession(req)
	s.Set("nom", "james")
	fmt.Fprintln(w, "OK")
}

func read(w http.ResponseWriter, r *http.Request) {
	s := sessions.GetSession(r)
	if s.Get("session_test_key") != "James" {
		log.Println("Failed to read")
	}
	fmt.Fprintf(w, "OK")
}

func delete(w http.ResponseWriter, r *http.Request) {
	s := sessions.GetSession(r)
	s.Set("session_key", "Jamie")
	s.Delete("session_key")
	fmt.Fprintf(w, "OK")
}

func check(w http.ResponseWriter, r *http.Request) {
	s := sessions.GetSession(r)
	if s.Get("session_key") == "Jamie" {
		log.Println("Failed to delete")
	}
	fmt.Fprintf(w, "OK")
}
