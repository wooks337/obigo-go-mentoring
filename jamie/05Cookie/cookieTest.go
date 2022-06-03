package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", CookieHandler())
}

func CookieHandler() http.Handler {
	m := mux.NewRouter()
	m.HandleFunc("/", create)
	m.HandleFunc("/readall", readAll)
	m.HandleFunc("/read", read)
	m.HandleFunc("/update", update)
	m.HandleFunc("/delete", delete)
	return m
}

//쿠키 (여러개) 생성
func create(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "cookiename",
		Value: "cookievalue",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "Hello",
		Value: "imvalue",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "Bonjour",
		Value: "jesuisvalueraussi",
	})

	fmt.Fprintln(w, "go to F12/application/cookie")
}

//모든 쿠키 읽기
func readAll(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()

	for _, ck := range cookies {
		fmt.Fprintln(w, ck)
	}
}

//쿠키 1개 읽기
func read(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("cookiename")

	if err == http.ErrNoCookie {
		fmt.Println("cookiename 이라는 쿠키명은 없습니다.")
	}
	fmt.Fprintln(w, "Your Cookie: ", cookie)
}

//쿠키 삭제
func delete(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("cookiename")

	//쿠키 바로 삭제
	cookie.MaxAge = -1
	//쿠키 10초뒤에 삭제
	//cookie.MaxAge = 10
	//Expires
	//expiration := time.Now().Add(time.Hour)
	//cookie.Expires = expiration
	http.SetCookie(w, cookie)
}

//쿠키 수정
func update(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("cookiename")
	cookie.Value = "1111"
	//cookie.Name = "Names" --- 이름은 수정 안됨
	cookie.Path = "/"
	cookie.Domain = "localhost"
	http.SetCookie(w, cookie)
}
