package main

import "fmt"

/*
1) Greeter에 사용되는 Message
2) Message를 담고있는 Greeter
3) Greeter가 Message로 환영하는 Event
*/

type Message string

func NewMessage() Message {
	return Message("Hi there!")
}

type Greeter struct {
	message Message
}

//Greeter 생성자
func NewGreeter(m Message) Greeter {
	return Greeter{message: m}
}

//Greet 메서드
func (g *Greeter) Greet() Message {
	return g.message
}

type Event struct {
	greeter Greeter
}

//Event 생성자
func NewEvent(g Greeter) Event {
	return Event{greeter: g}
}

func (e Event) Start() {
	msg := e.greeter.Greet()
	fmt.Println(msg)
}

func main() {
	//message := NewMessage()
	//greeter := NewGreeter(message)
	//event := NewEvent(greeter)
	//
	//event.Start()

	//Pour wire package
	e := InitializeEvent()

	e.Start()
}
