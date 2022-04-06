//회원가입 시 암호 길이를 검사하는 예제
package main

import "fmt"

//에러 구조체 선언
//- 암호 길이 관련 에러 정보를 담기 위한 구조체 선언
type PasswordError struct {
	Len        int
	RequireLen int
}

//Error() 메소드 선언
//- Error()메소드는 PasswordError 구조체에 속함
//- Error()메소드는 error 인터페이스에 속하는 메소드
//- 따라서 PasswordError 구조체는 error 인터페이스로 사용가능
func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

//RegisterAccount() 함수 - 매개변수 : name, password - 반환타입 : error
//- 반환타입이 error이지만 PasswordError 구조체가 error 인터페이스로 사용가능
//- 따라서 RegisterAccount() 함수는 PasswordError로 에러 반환
func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8} //입력 길이와 필요 길이 반환
	}
	return nil
}

//RegisterAccount()함수로 Id와 Pw 입력
//if 에러가 nil이 아니고
//	if err->PasswordError 타입변환이 ok이면, Printf()
//if 에러가 nil이면(비밀번호가 잘 입력 되었으면), -> '회원 가입되었습니다' 출력
func main() {
	err := RegisterAccount("myID", "myPw") // ID, Pw 입력
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok { // ; ok { 는 아래 코드를 단순화 한 것!
			fmt.Printf("%v Len:%d RequireLen:%d\n",
				errInfo, errInfo.Len, errInfo.RequireLen)
		}
		//if errInfo, ok := err.(PasswordError) {
		//	if ok {
		//		fmt.Printf("%v Len:%d RequireLen:%d\n",
		//			errInfo, errInfo.Len, errInfo.RequireLen)
		//	}
		//}
	} else {
		fmt.Println("회원 가입됐습니다.")
	}
}

//Type Conversion과 Type Assertion의 차이 알아두기!
