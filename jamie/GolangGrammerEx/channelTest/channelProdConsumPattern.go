package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup      //waitGroup 객체 생성
var startTime = time.Now() //시작 시간

func main() {
	//채널 생성 - Car 구조체의 메모리 주소를 메세지 타입으로 가짐
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n") //생산 가동

	wg.Add(3) //고루틴 작업 3개

	//고루틴 생성
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait() //전체 단계가 종료될 때까지 대기
	fmt.Println("Close the factory")
}

//차체 생산
func MakeBody(tireCh chan *Car) {
	//시간 객체 생성
	tick := time.Tick(time.Second)        //1초
	after := time.After(10 * time.Second) //10초

	for {
		select {
		case <-tick: //1초 간격
			//Make a Body
			car := &Car{} //Car 구조체 초기화
			car.Body = "Sport car"
			tireCh <- car //tireCh 채널에 car.Body("sport car") 데이터 주입 = 1초간격으로!
		case <-after: //10초 뒤
			close(tireCh) //10초뒤 tireCh 종료!
			wg.Done()     //차체 생산 루틴 완료
			return
		}
	}
}

//바퀴 설치
func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh { //tireCh가 주입받은 car 데이터를 읽어서 바퀴를 설치!
		//Make a Body
		time.Sleep(time.Second) //1초 대기
		car.Tire = "Winter tire"
		paintCh <- car //paintCh 채널에 car.Tire("Winter tire") 데이터 주입
	}
	wg.Done() //바퀴 설치 루틴 완료
	close(paintCh)
}

//도색
func PaintCar(paintCh chan *Car) {
	for car := range paintCh { //paintCh가 주입받은 car 데이터를 읽어서 도색!
		//Make a Body
		time.Sleep(time.Second) //1초 대기
		car.Color = "Red"
		duration := time.Now().Sub(startTime)                                                          //경과 시간(현재시각-시작시간)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color) //경과 시간 및 데이터 출력
	}
	wg.Done() // 도색 루틴 완료
}
