package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Message string

//Message 생성자
func NewMessage() Message {
	return Message("Hi there!")
}

//Greeter 구조체
type Greeter struct {
	message Message
	grumpy  bool
}

//Greeter 생성자
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{message: m, grumpy: grumpy}
}

//Greet 메서드
func (g *Greeter) Greet() Message {
	if g.grumpy {
		return Message("Go Away...Im not in the mood")
	}
	return g.message
}

//Event 구조체
type Event struct {
	greeter Greeter
}

//Event 생성자 - 에러 리턴 하도록 변경
func NewEvent(g Greeter) (Event, error) {
	if g.grumpy {
		return Event{}, errors.New("이벤트 생성 실패 : event Greeter is grumpy :(")
	}
	return Event{greeter: g}, nil
}

//Event Start() 메소드
func (e Event) Start() {
	msg := e.greeter.Greet()
	fmt.Println(msg)
}

func main() {
	e, err := InitializeEvent()
	if err != nil {
		fmt.Printf("이벤트 생성 실패 : %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
