package main

import "fmt"

func main() {

	err := RegisterAccount("myId", "pwd")
	if err != nil {
		if passwordError, ok := err.(PasswordError); ok { //인터페이스 변환
			fmt.Println(passwordError)
		}
	}
}

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다"
}

func RegisterAccount(name, password string) error { //error 인터페이스
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}
	return nil
}
