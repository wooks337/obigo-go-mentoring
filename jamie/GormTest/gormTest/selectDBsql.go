package main

import (
	"database/sql"
	"fmt"
	_ "gorm.io/driver/mysql"
	"log"
)

//======Localhost======
func main() {
	dsn := "root:jamiekim@(localhost:3306)/sakila?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	} else {
		defer db.Close()
	}
	//======TeamServer======
	//func main() {
	//	dsn := "root:root@(10.28.3.180:3307)/sakila?charset=utf8mb4&parseTime=True&loc=Local"
	//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//		Logger: logger.Default.LogMode(logger.Info),
	//	})
	//	if err != nil {
	//		panic(err)
	//	}

	//======단일 레코드 조회======
	var first_name string
	var last_name string
	var actor_id int
	err = db.QueryRow("SELECT first_name FROM actor LIMIT 1").Scan(&first_name)
	if err != nil {
		log.Fatalln("err 발생 : ", err)
	}
	fmt.Println(first_name)

	//======복수 레코드 조회======
	//======Placeholder ? 를 이용하여 데이터 대입
	rows, err := db.Query("SELECT actor_id, first_name, last_name FROM actor where actor_id < ?", 11)
	if err != nil {
		log.Fatalln("err 발생 : ", err)
	}
	defer rows.Close()

	for rows.Next() { //======row에서 다음 row로 이동하기 위해 Next()메소드 이용
		err := rows.Scan(&actor_id, &first_name, &last_name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(actor_id, first_name, last_name)
	}
}
