package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		stdin.ReadString('\n')
	}
	return num, err
}

func main() {

	rand.Seed(time.Now().UnixNano()) //현재 시간값(time.Now())을 랜덤 시드로 설정 -> seed 입력값이 int64이므로 UnixNano 메서드로 현재시각을 int64 타입으로 변경
	n := rand.Intn(100)              //0 ~ (n-1)사이의 값을 랜덤 생성하는 함수

	cnt := 1

	for {
		fmt.Print("1~99 사이의 숫자값를 입력하세요: ")
		user, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요!!")
		} else {
			if user > n {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if user < n {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수: ", cnt)
				break
			}
			cnt++
		}
	}
}
