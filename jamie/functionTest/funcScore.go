package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "님의 평균점수는 ", avg, "입니다.")
}

func main() {
	PrintAvgScore("김일등", 80, 74, 95)
	PrintAvgScore("송이등", 80, 50, 99)
	PrintAvgScore("박삼등", 25, 35, 40)
}
