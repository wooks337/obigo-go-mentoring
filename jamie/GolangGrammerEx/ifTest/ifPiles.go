//중첩 If문 예시
package main

import "fmt"

//부자인 친구가 있는가 - 무조건 true 반환
func HasRichFriend() bool {
	return true
}

// 같이 간 친구의 숫자 - 무조건 3 반환
func FriendCnt() int {
	return 3
}

func main() {
	price := 35000

	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("신발끈!!")
		} else {
			fmt.Println("나눠내자!")
		}
	} else if price >= 30000 && price <= 50000 {
		if FriendCnt() > 3 {
			fmt.Println("신발끈!!")
		} else {
			fmt.Println("나눠내자!")
		}
	} else {
		fmt.Println("내가낸다")
	}

}
