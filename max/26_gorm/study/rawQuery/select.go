package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_6_gorm "goMod"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(10.28.3.180:3307)/sakila?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("err발생 : ", err)
	} else {
		defer db.Close()
	}

	//========단일 데이터 조회=======
	//var title string
	//err = db.QueryRow("SELECT title FROM film LIMIT 1").Scan(&title)
	//if err != nil {
	//	fmt.Println("err 발생 : ", err)
	//}
	//fmt.Println(title)
	//
	//title = db.QueryRow("SELECT title FROM film WHERE film_id=-1").Scan(&title)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		fmt.Println("No Rows")
	//	} else {
	//		fmt.Println("err 발생 : ", err)
	//	}
	//} else {
	//	fmt.Println(title)
	//}

	//========다수 데이터 조회=======
	//rows, err := db.Query("SELECT title FROM film LIMIT 10")
	//if err != nil {
	//	fmt.Println("err발생 : ", err)
	//} else {
	//	for rows.Next() {
	//		var title string
	//		rows.Scan(&title)
	//		fmt.Println(title)
	//	}
	//}

	rows, err := db.Query("SELECT film_id, title FROM film LIMIT 10")
	if err != nil {
		fmt.Println("err 발생 : ", err)
	} else {
		for rows.Next() {
			f := _6_gorm.Film{}
			rows.Scan(&f.Film_id, &f.Title)
			fmt.Println(f.Film_id, ", ", f.Title)
		}
	}

}
