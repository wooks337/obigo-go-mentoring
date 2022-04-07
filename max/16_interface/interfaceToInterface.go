package main

func main() {

	f := file{}         //구조체 생성
	r := reader(f)      //구조체 -> 인터페이스
	w, ok := r.(writer) //인터페이스 -> 인터페이스, 그러나 write() 메서드 없어서 실패
	if ok {
		println("변환 성공", w.write)
	} else {
		println("변환 실패")
	}
}

type reader interface {
	read()
}

type writer interface {
	write()
}

type file struct {
}

func (f file) read() {
}
