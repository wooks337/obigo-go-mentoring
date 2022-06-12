package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"jamie/domain"
)

//mysql 서버 연결 함수
func ConnectDB() (*gorm.DB, error) {
	dsn := "root:jamiekim@(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:root@tcp(10.28.3.180:3307)/Jamie?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db, err
}

//db 데이터 저장 함수
func SignUp(db *gorm.DB, user domain.User) error {
	res := db.Create(&user)
	return res.Error
}

//아이디 중복 체크 함수
func IDCheck(db *gorm.DB, userid string) bool {
	findID := domain.User{}
	res := db.Model(&domain.User{}).First(&findID, "userid = ?", userid)
	if res.Error != nil {
		return true
	} else {
		return false
	}
}

//비밀번호 암호화 함수
//https://bourbonkk.tistory.com/64
//https://jeong-dev-blog.tistory.com/2
//pwHash, _ := bcrypt.GenerateFromPassword([]byte(), bcrypt.DefaultCost)
//[]byte 자료형의 해시 반환 -> 해시 반환값을 string 변환 후 DB 저장
func HashPassword(password string) (string, error) {
	pw := []byte(password)

	pwHash, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", pwHash), nil
}

//userid로 회원정보 찾기 함수
func FindUserByUserid(db *gorm.DB, userid string) (domain.User, error) {
	//User 구조체에 회원조회 정보 담아서 에러랑 같이 반환하기
	findUser := domain.User{}
	res := db.Model(&domain.User{}).First(&findUser, "userid = ?", userid)
	return findUser, res.Error
}

//로그인 시 비밀번호 일치 확인 함수
func CheckPasswordHash(hashVal, userPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw))
	if err != nil {
		return false
	} else {
		return true
	}
}
