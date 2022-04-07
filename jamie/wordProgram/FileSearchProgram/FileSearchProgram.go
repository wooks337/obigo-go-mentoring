package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//찾은 line 정보
type LineInfo struct { //찾은 결과 정보
	lineNo int
	line   string
}

//파일 내 라인 정보
type FindInfo struct {
	filename string
	lines    []LineInfo //line들 정보가 담긴 LineInfo 구조체 배열 필드
}

func main() {
	//실행 인수 개수 확인
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행인수가 필요합니다.")
		return
	}
	//실행 인수 가져오기
	word := os.Args[1]        //찾으려는 단어 인수
	files := os.Args[2:]      //찾으려는 파일명 인수
	findInfos := []FindInfo{} //빈 슬라이스 객체 생성 - 파일 내 찾은 단어 저장
	for _, path := range files {
		//파일 찾기
		//FileWordsAllFiles 함수() 호출 -> 파일 내 단어 찾고 findInfos 슬라이스 내에 추가
		findInfos = append(findInfos, FindWordsAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos { //findInfos 배열 내에서 findInfo 요소(filename) 조회 및 출력
		fmt.Println(findInfo.filename)
		fmt.Println("----------------------------")
		for _, lineInfo := range findInfo.lines { //findInfo.lines(=[]LineInfo)의 배열 내에서 lineInfo 요소(lineNo, line) 조회 및 출력
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("-------------------------")
		fmt.Println() // 줄띄우기
	}
}

//파일 목록 가져오는 함수
func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path) //파일 경로에 해당하는 파일 리스트를 []string 타입으로 변환
}

//파일 목록 출력하는 함수
func FindWordsAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := GetFileList(path) //GetFileList() 함수로 가져온 파일 리스트를 filelist 객체에 주입
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err:", err)
		return findInfos
	}
	for _, filename := range filelist { //filelist내 각 파일별 단어 검색 및 findInfos 슬라이스 객체에 저장 - FindWordInFile() 함수 이용
		findInfos = append(findInfos, FindWordInFile(word, filename))
	}
	return findInfos
}

//파일 내 단어 검색하는 함수
func FindWordInFile(word, filename string) FindInfo {
	findInfo := FindInfo{filename, []LineInfo{}} //빈 슬라이스 객체 - 파일 내 단어가 있는 라인 정보 저장

	file, err := os.Open(filename) //파일 열기
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		return findInfo
	}
	defer file.Close() //함수 종료 전 파일 닫기

	lineNo := 1                       //1번 라인부터
	scanner := bufio.NewScanner(file) //파일 내용 한줄씩 읽기
	for scanner.Scan() {              //다음줄을 읽어옴
		line := scanner.Text() //읽어온 라인을 문자열로 반환

		if strings.Contains(line, word) { //line 안에 word가 포함되어 있다면 findInfo.lines에 추가 및 저장
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++ //라인은 1씩 증가
	}
	return findInfo
}
