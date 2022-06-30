package main

import (
	"11Cron/01_DBConnect/service"
	"fmt"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func main() {
	//DB연결
	db, err := service.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't connect to DB!")
	}
	log.Println("Connected to Database...")
	defer func() {
		d, _ := db.DB()
		d.Close()
		log.Println("Database Closed...")
	}()

	////테이블 생성
	//db.AutoMigrate(&domain.TestInfo{})

	////예시 데이터 입력

	Cron()
}

func Cron() {

	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Minutes().Do(service.UpdateStatus)
	s.Every(1).Minute().Do(service.UpdateMin)
	s.Every(30).Seconds().Do(service.UpdateSec)

	s.StartAsync()
	for {
		time.Sleep(time.Second)
	}
}
