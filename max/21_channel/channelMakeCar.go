package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Println("Start Factory")
	wg.Add(3)
	go MakeBody(tireCh)             //차체 생산
	go InstallTire(tireCh, paintCh) //차체 생산 확인 후 바퀴설치
	go PaintCar(paintCh)            //바퀴 설치 확인 후 색칠
	wg.Wait()
	fmt.Println("Close Factory")
}

type Car struct {
	Body  string
	Tire  string
	Color string
}

func MakeBody(tireCh chan *Car) { //차체 생산
	tick := time.Tick(time.Second) //매초마다 실행
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports Car"
			tireCh <- car //차체 완성 후 tire채널에 넣음
		case <-after:
			close(tireCh) //채널을 닫아 고루틴릭 상태에서 벗어남
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) { //바퀴설치

	for car := range tireCh { //차체 생산 대기 후 루프
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintCh <- car //paint채널에 넣음
	}
	wg.Done()
	close(paintCh) //채널을 닫아 고루틴릭 상태에서 벗어남
}

func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Println("Complete Car, ", duration.Seconds(), ", ", car.Body, ", ",
			car.Tire, ", ", car.Color)
	}
	wg.Done()
}
