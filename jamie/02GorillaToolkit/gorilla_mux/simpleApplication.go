package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// staticPath와 indexPath를 필드로 가지는 spaHandler 구조체
// spaHandler는 http.Handler 인터페이스를 가져 HTTP요청에 응답할 수 있다
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP는 URL 경로를 검사하여 SPA 처리기의 정적 디렉터리 내에서 파일을 찾아서 전달한다.
// 파일이 존재하지 않을 경우, SPA 핸들러의 인덱스 경로에 있는 파일이 제공된다.
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 절대 경로 가져오기
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// 절대경로를 못찾으면 400 bad request 출력
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 경로 앞에 정적 디렉토리 추가
	path = filepath.Join(h.staticPath, path)

	// 해당 경로 내 파일 존재 여부 확인
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// 파일이 없으면 index.html 전달
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// http.FileServer로 정적 디렉토리 전달
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	router := mux.NewRouter()
	// API 핸들러
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	spa := spaHandler{staticPath: "build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
