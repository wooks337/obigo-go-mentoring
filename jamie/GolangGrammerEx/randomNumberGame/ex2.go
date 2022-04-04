package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//슬롯머신 게임
/*
1000원으로 시작
1~5 사이의 값을 입력받음
1~5사이의 랜덤값 선택
if, 입력값과 난수가 일치하면 +500원, 축하메세지와 잔액 표시
if, 불일치하면 -100원, 아쉬움메세지와 잔액표시
반복
가진 돈이 0이하 || 5000원 이상 -> 게임 종료
*/

//상수값으로 기본값 설정
const (
	//초기잔액
	Balance = 1000
	//이긴금액
	EarnPoint = 500
	//진금액
	LosePoint = 100
	//승리종료금액
	VictoryPoint = 5000
	//게임오버포인트
	GameOverPoint = 0
)

var stdin = bufio.NewReader(os.Stdin)

func InputNum() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}
func main() {
	rand.Seed(time.Now().UnixNano())
	balance := Balance

	for {
		fmt.Println("1~5사이의 값을 입력하세요: ")
		n, err := InputNum()
		if err != nil {
			fmt.Println("숫자만 입력하세요!")
		} else if n < 1 || n > 5 {
			fmt.Println("1~5사이의 숫자만 입력하세요!")
		} else {
			ranNum := rand.Intn(5) + 1
			if ranNum == n {
				balance += EarnPoint
				fmt.Println("축하합니다! 현재 잔액은 ", balance, "원 입니다.")
				if balance > VictoryPoint {
					fmt.Println("~~~게임 승리~~~")
					break
				}
			} else {
				balance -= LosePoint
				fmt.Println("아쉽네요! 현재 잔액은", balance, "원 입니다.")
				if balance <= GameOverPoint {
					fmt.Println("~~~게임 오버~~~")
					break
				}
			}
		}
	}

}
