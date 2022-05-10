package main

import (
	"fmt"
	"goMod/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	dsn := "root:root@tcp(10.28.3.180:3307)/max?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //-- 모든 SQL 실행문 로그로 확인
	})

	if err != nil {
		err := fmt.Errorf("연결실패 : %v", err)
		panic(err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.Delete(&domain.Student{}) //조건없어서 오류

	//db.Where("1=1").Delete(&domain.Student{})
	//db.Exec("DELETE FROM student")
	//db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.Student{})

}
