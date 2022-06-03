package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//실행 인수 개수 확인
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다.")
		return
	}

	//실행 인수 가져오기기
	word := os.Args[1]
	files := os.Args[2:]
	fmt.Println("찾으려는 단어: ", word)
	PrintAllFiles(files)
}

//파일 목록 가져오는 함수
func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path) //파일 경로에 해당하는 파일 리스트를 []string 타입으로 변환
}

//파일 목록 출력하는 함수
func PrintAllFiles(files []string) {
	for _, path := range files {
		filelist, err := GetFileList(path) //파일 목록 가져오기
		if err != nil {
			fmt.Println("파일을 찾을 수 없습니다. err:", err)
			return
		}
		fmt.Println("찾으려는 파일 리스트")
		for _, name := range filelist {
			fmt.Println(name)
		}
	}
}
