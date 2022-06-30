package service

import (
	"11Cron/01_DBConnect/domain"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(10.28.3.180:3307)/Jamie?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //-- 모든 SQL 실행문 로그로 확인
	})
	return db, err
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
