package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex //패키지 전역변수 뮤텍스 선언

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()         //뮤텍스 획득
	defer mutex.Unlock() //프로그램 종료 전에 defer로 mutex unlock
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value :%d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0} //잔고가 0 인 계좌
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
