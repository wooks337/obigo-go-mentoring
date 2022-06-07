package main

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"time"
)

var rd *render.Render

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

	router.HandleFunc("/create/{name}", createCookieHandler).Methods("GET")
	router.HandleFunc("/read", readAllCookieHandler).Methods("GET")
	router.HandleFunc("/read/{name}", readCookieHandler).Methods("GET")
	router.HandleFunc("/remove/{name}", removeCookieHandler).Methods("GET")
	return router
}

func createCookieHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	name := vars["name"]

	http.SetCookie(w, &http.Cookie{
		Name:       name,
		Value:      "value" + name,
		Path:       "/",
		Domain:     "",
		Expires:    time.Time{}, //Deprecated
		RawExpires: "",
		MaxAge:     600, //초
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	})
	http.Redirect(w, req, "/read", http.StatusCreated)
	//rd.JSON(w, http.StatusOK, Success{true}) //JSON 변환
}

func readAllCookieHandler(w http.ResponseWriter, req *http.Request) {

	cookieList := req.Cookies()

	rd.JSON(w, http.StatusOK, cookieList)
}

func readCookieHandler(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	name := vars["name"]

	cookie, err := req.Cookie(name)
	if err != nil {
		rd.JSON(w, http.StatusNoContent, err)
		return
	}

	rd.JSON(w, http.StatusOK, cookie)
}

func removeCookieHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]

	cookie, err := req.Cookie(name)
	if err != nil {
		rd.JSON(w, http.StatusNoContent, err)
		return
	}

	cookie.MaxAge = -1
	cookie.Path = "/"
	http.SetCookie(w, cookie)
	rd.JSON(w, http.StatusOK, cookie)
}

type Success struct {
	Success bool `json:"success"`
}
