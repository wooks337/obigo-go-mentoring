package main

import (
	"11Cron/DBConnect/domain"
	"fmt"
	"github.com/go-co-op/gocron"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var db *gorm.DB

func main() {
	//DB연결
	db, err := ConnectDB()
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
	//d1 := domain.TestInfo{Name: "Ian"}
	//d2 := domain.TestInfo{Name: "Park"}
	//d3 := domain.TestInfo{Name: "Kim"}
	//d4 := domain.TestInfo{Name: "Lilly"}
	//d5 := domain.TestInfo{Name: "Robert"}
	//db.Create(&d1)
	//db.Create(&d2)
	//db.Create(&d3)
	//db.Create(&d4)
	//db.Create(&d5)

	Cron()
}

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(10.28.3.180:3307)/Jamie?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //-- 모든 SQL 실행문 로그로 확인
	})
	return db, err
}

func Cron() {

	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Minutes().Do(UpdateStatus)
	s.Every(1).Minute().Do(UpdateMin)
	s.Every(30).Seconds().Do(UpdateSec)

	s.StartAsync()
	for {
		time.Sleep(time.Second)
	}
}

func UpdateStatus() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	db.Model(domain.TestInfo{}).Where("status = ?", "off").Updates(domain.TestInfo{Status: "on"})
	log.Println("5분 단위 업데이트")
}

func UpdateMin() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	res := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(domain.TestInfo{}).Updates(map[string]interface{}{"Min": gorm.Expr("Min + ?", 1)})
	fmt.Println("1분에 바뀐 데이터 갯수", res.RowsAffected)
	log.Println("1분 단위 업데이트")
}
func UpdateSec() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	res := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(domain.TestInfo{}).Where("1=1").Updates(map[string]interface{}{"Sec": gorm.Expr("Sec + ?", 30)})
	fmt.Println("30초에 바뀐 데이터 갯수", res.RowsAffected)
	log.Println("30초 단위 업데이트")
}
