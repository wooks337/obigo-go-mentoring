package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/fedex"
	"github.com/tuckersGo/musthaveGo/ch20/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	// 우체국 전송 객체
	sender := &koreaPost.PostSender{} // *koreaPost.PostSender 타입
	SendBook("어린 왕자", sender)         // 타입 오류
	SendBook("그리스인 조르바", sender)
}
