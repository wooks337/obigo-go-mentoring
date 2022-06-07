package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

var rd *render.Render

var key = []byte("super-secret-key")

//var key = make([]byte, 32)

var store = sessions.NewCookieStore(key)

func main() {
	rd = render.New()
	mux := makeWebHandler()
	n := negroni.Classic() //각 요청이 올 때마다 터미널에 로그가 찍힘
	n.UseHandler(mux)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}

func makeWebHandler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/create/{name}", createSessionHandler).Methods("GET")
	router.HandleFunc("/read/{name}", readSessionHandler).Methods("GET")
	router.HandleFunc("/remove/{name}", removeSessionHandler).Methods("GET")
	return router
}

func createSessionHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	name := vars["name"]

	session, _ := store.Get(req, name)    //세션 생성, 존재하지 않을경우 생성
	session.Values[name] = name + "value" //새션스토어에 데이터 저장
	session.Options.MaxAge = 1 * 60
	err := session.Save(req, w)
	if err != nil {
		rd.JSON(w, http.StatusNoContent, err)
		return
	}
	http.Redirect(w, req, "/read/"+name, http.StatusCreated)
}

func readSessionHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	name := vars["name"]

	session, _ := store.Get(req, name)
	value := session.Values[name]
	rd.JSON(w, http.StatusOK, value)
}

func removeSessionHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	name := vars["name"]

	session, _ := store.Get(req, name)
	session.Options.MaxAge = -1
	err := session.Save(req, w)
	if err != nil {
		rd.JSON(w, http.StatusNoContent, err)
		return
	}

	rd.JSON(w, http.StatusOK, Success{true})
}

type Success struct {
	Success bool `json:"success"`
}
