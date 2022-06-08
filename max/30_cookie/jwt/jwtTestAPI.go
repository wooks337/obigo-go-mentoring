package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	//"github.com/dgrijalva/jwt-go/v4"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"strconv"
	"time"
)

var rd *render.Render
var signKey = []byte("secret")

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

	router.HandleFunc("/create", createJWTHandler).Methods("GET")
	router.HandleFunc("/read", readJWTHandler).Methods("GET")
	router.HandleFunc("/remove", removeJWTHandler).Methods("GET")
	return router
}

func createJWTHandler(w http.ResponseWriter, req *http.Request) {

	token, err := createJWT(req)
	if err != nil {
		rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(token)
	http.SetCookie(w, &http.Cookie{
		Name:       "jwt",
		Value:      token,
		Path:       "/",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     20 * 60,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	})
	http.Redirect(w, req, "/read", http.StatusCreated)
	//rd.JSON(w, http.StatusOK, Success{true}) //JSON 변환
}

func createJWT(req *http.Request) (string, error) {
	query := req.URL.Query()
	name := query.Get("name")
	age, _ := strconv.Atoi(query.Get("age"))

	//payload 셋팅
	claims := userClaims{
		Name: name,
		Age:  age,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
			//ExpiresAt: jwt.At(time.Now().Add(time.Minute * 1)), //만료시간 설정
		},
	}
	//헤더 및 페이로드 생성
	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	//서명
	signedString, err := aToken.SignedString(signKey)

	return signedString, err
}

func readJWTHandler(w http.ResponseWriter, req *http.Request) {

	res := make(map[string]interface{})

	cookie, err := req.Cookie("jwt")
	if err != nil {
		rd.JSON(w, http.StatusUnauthorized, err.Error())
		return
	}

	res["CookieName"] = cookie.Name
	res["CookieValue"] = cookie.Value

	var claims userClaims
	//var mapClaims jwt.MapClaims
	tok, err := jwt.ParseWithClaims(cookie.Value, &claims, func(token *jwt.Token) (interface{}, error) {
		return signKey, err
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors == jwt.ValidationErrorExpired {
				rd.JSON(w, http.StatusUnauthorized, "유효기간 만료")
				return
			}
		}
		rd.JSON(w, http.StatusUnauthorized, err)
		return
	}

	fmt.Println(tok.Valid)
	res["claims"] = claims

	rd.JSON(w, http.StatusOK, res)
}

func removeJWTHandler(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("jwt")
	if err != nil {
		rd.JSON(w, http.StatusNoContent, err.Error())
		return
	}

	cookie.MaxAge = -1
	cookie.Path = "/"
	http.SetCookie(w, cookie)
	rd.JSON(w, http.StatusOK, cookie)
}

func updateJWT(claims userClaims) (string, error) {

	//헤더 및 페이로드 생성
	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	//서명
	signedString, err := aToken.SignedString(signKey)

	return signedString, err
}

type success struct {
	Success bool `json:"success"`
}

type userClaims struct {
	Name string
	Age  int
	jwt.StandardClaims
}
