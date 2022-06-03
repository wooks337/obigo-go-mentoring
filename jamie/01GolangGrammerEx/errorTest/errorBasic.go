package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename) //파일 열기
	if err != nil {
		return "", err //파일 열고 에러나면 에러 반환
	}
	defer file.Close()             //ReadFile()함수 종료 직전 파일 닫기
	rd := bufio.NewReader(file)    // 파일 내용 읽기
	line, _ := rd.ReadString('\n') //한줄의 line 읽기, error는 문자열이 \n으로 끝나지 않을 경우 에러 반환 -> _로 에러 무시
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename) //파일 생성 (파일 핸들과 에러 반환)
	if err != nil {                  //파일 만들고 에러나면 에러 반환
		return err
	}
	defer file.Close()                //함수 종료 직전 파일 닫기
	_, err = fmt.Fprintln(file, line) //파일에 line 문자열 쓰기(작성 문자열 길이와 에러 반환 -> 길이 _ 로 무시)
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename) //파일 읽기 시도
	if err != nil {
		err = WriteFile(filename, "This is WriteFile") //에러 발생 시 WriteFile()함수 호출하여 파일 생성
		if err != nil {                                //파일 생성시에도 에러 발생 시 하기 메세지 출력 후 프로그램 종료
			fmt.Println("파일 생성에 실패하였습니다.", err)
			return
		}
		line, err = ReadFile(filename) //파일 생성 성공 시, 다시 읽기 시도
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}
	}
	fmt.Println("파일내용:", line) //파일내용 출력
}
