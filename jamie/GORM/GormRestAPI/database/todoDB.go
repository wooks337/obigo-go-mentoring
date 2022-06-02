package database

//할일 정보 담는 구조체
type Todo struct {
	ID        int    `json:"id,omitempty"` //json 포맷 변환 옵션(항목 이름 : id, 생략 가능)
	Name      string `json:"name"`
	Completed bool   `json:"completed,omitempty"`
}
