package domain

import "gorm.io/gorm"

//회원가입 DB 저장용(pw암호화)
type User struct {
	gorm.Model
	UserID   string `gorm:"unique;not null" json:"userid"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null" json:"email"`
}

//회원가입 시 입력용 : DB로 구조화
type JoinUser struct {
	UserID   string `json:"userid"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

//로그인용 : DB로 Select함
type LoginUser struct {
	UserID   string `json:"userid"`
	Password string `json:"password"`
}

//회원 페이지용
type InfoUser struct {
	ID     uint   `json:"id"`
	UserID string `json:"userid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

//구글로그인 유저용
type GoogleUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
