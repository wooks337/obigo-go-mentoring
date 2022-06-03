//실행할 때마다 랜덤한 숫자 출력
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //시간값을 랜덤 시드로 설정

	n := rand.Intn(100)
	fmt.Println(n)

}
