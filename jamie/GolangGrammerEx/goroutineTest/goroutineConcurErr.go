package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}    //잔고 0원 통장; Account 구조체 포인터
	wg.Add(10)                //waitGroup 객체 생성; 작업 개수 지정
	for i := 0; i < 10; i++ { //for 반복문으로 고루틴 10개 생성
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done() //작업 완료
		}() //함수리터럴 func() 호출
	}
	wg.Wait() //모든 작업 종료시 까지 대기
}
func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value :%d", account.Balance))
	}
	account.Balance += 1000      //1000원 입금
	time.Sleep(time.Millisecond) //잠시 대기
	account.Balance -= 1000      //1000원 출금
}

//출력
//잔고가 0미만으로 내려가 panic 발생
//account.Balance += 1000는 Balance 값을 읽고 1000을 더해 Balance를 저장하는 두 단계
//첫 번째 단계가 끝나기 전, 다른 고루틴이 첫 번째 단계를 수행할 경우,
//두 고루틴이 똑같은 값을 읽어 각각 입금을 하지만 한번 입금한 효과가 남 -> 각가 출금이 이루어질 경우 마이너스 잔고가 된다.
