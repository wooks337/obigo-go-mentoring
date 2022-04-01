package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/fedex"
	"github.com/tuckersGo/musthaveGo/ch20/koreaPost"
)

// Sender 인터페이스
type Sender interface {
	Send(parcel string)
}

// 함수에 Sender 인터페이스 입력받기
func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	// 우체국 전송 객체
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("어린 왕자", koreaPostSender)
	SendBook("그리스인 조르바", koreaPostSender)

	// Fedex 전송 객체
	fedexSender := &fedex.FedexSender{}
	SendBook("어린 왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)
}
