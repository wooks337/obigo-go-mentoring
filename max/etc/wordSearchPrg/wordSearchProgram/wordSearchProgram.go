package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다.")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	findInfos := []FindInfo{}

	basicStart := time.Now()

	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("-----------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("  ", lineInfo.lineNo, "  ", lineInfo.line)
		}
		fmt.Println("-----------------------")
		fmt.Println()
	}
	fmt.Println("일반 소요시간 : ", time.Now().Sub(basicStart).Seconds())

	goRoutineStart := time.Now()

	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFilesGoRoutine(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("-----------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("  ", lineInfo.lineNo, "  ", lineInfo.line)
		}
		fmt.Println("-----------------------")
		fmt.Println()
	}
	fmt.Println("고루틴 소요시간 : ", time.Now().Sub(goRoutineStart).Seconds())
}

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

//고루틴으로 실행하기
//path에 있는 모든 파일들을 찾음
func FindWordInAllFilesGoRoutine(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	fileList, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다 err : ", err)
		return findInfos
	}

	ch := make(chan FindInfo)
	cnt := len(fileList)
	recvCnt := 0

	for _, fileName := range fileList {
		go FindWordInFileGoRoutine(word, fileName, ch)
	}

	for findInfo := range ch {
		findInfos = append(findInfos, findInfo)
		recvCnt++
		if recvCnt == cnt {
			//close(ch)
			break
		}
	}
	return findInfos
}

//고루틴 사용
//한개의 파일에서 한줄씩 단어를 검색하여 FindInfo에 저장
func FindWordInFileGoRoutine(word, fileName string, ch chan FindInfo) {

	findInfo := FindInfo{fileName, []LineInfo{}}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", fileName)
		ch <- findInfo
		return
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo
}

//path에 있는 모든 파일들을 찾음
func FindWordInAllFiles(word, path string) []FindInfo {

	findInfos := []FindInfo{}

	fileList, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다 err : ", err)
		return findInfos
	}

	for _, fileName := range fileList {
		findInfos = append(findInfos, FindWordInFile(word, fileName))
	}
	return findInfos
}

//한개의 파일에서 한줄씩 단어를 검색하여 FindInfo에 저장
func FindWordInFile(word, fileName string) FindInfo {

	findInfo := FindInfo{fileName, []LineInfo{}}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", fileName)
		return findInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	return findInfo
}

//해당 경로의 파일목록을 가져옴
func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}
