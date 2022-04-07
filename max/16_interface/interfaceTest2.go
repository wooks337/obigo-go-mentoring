package main

import "fmt"

func main() {

	k := korMail{}
	u := usaMail{}
	SendMailK("서울", k)
	SendMailU("newyork", u) //비슷한 기능이지만 따로 사용

	SendMail("서울", k)
	SendMail("newyork", u)
}

func SendMail(addr string, sender Sender) {
	sender.Send(addr)
}

type Sender interface {
	Send(addr string)
}

func SendMailK(addr string, mail korMail) {
	mail.Send(addr)
}
func SendMailU(addr string, mail usaMail) {
	mail.Send(addr)
}

type korMail struct {
}

func (k korMail) Send(addr string) {
	fmt.Println("주소는 ", addr, "입니다")
}

type usaMail struct {
}

func (u usaMail) Send(addr string) {
	fmt.Println("Address is ", addr)
}
