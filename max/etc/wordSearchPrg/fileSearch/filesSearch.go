package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
	실행 인수 읽고 파일 목록 가져오기
*/
func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) word ")
		return
	}

	find := os.Args
	word := os.Args[1]   //찾을 단어
	files := os.Args[2:] //파일경로

	fmt.Println("실행명령 : ", find)
	fmt.Println("찾으려는 단어 : ", word)
	PrintAllFiles(files)
}

func PrintAllFiles(files []string) {

	for _, path := range files {
		fileList, err := GetFileList(path)
		if err != nil {
			fmt.Println("파일을 찾을 수 없습니다. err : ", err)
			return
		}
		fmt.Println("찾으려는 파일 리스트")
		for _, name := range fileList {
			fmt.Println(name)
		}
	}

}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}
