package main

import "fmt"

var globalA = 5 // 함수 밖에서는 'var' 명시, 타입은 명시 안 해도 됨
const (
	c1 = 10 //첫 번째 값은 무조건 초기화해야 함
	c2
	c3
	c4 = "goorm" //다른 형 선언 가능
	c5
	c6 = iota //c8까지 index값 저장
	c7
	c8
	c9 = "earth"
	c10
	c11 = "End"
)

func main() {
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11)

	var a string = "goorm"
	fmt.Println(a)

	var b int = 10
	fmt.Println(b)

	var c = true
	fmt.Println(c)

	var d int
	fmt.Println(d)

	f := "short"
	fmt.Println(f)

	fmt.Println(globalA)

	var k, j int = 10, 11
	fmt.Println(k, j)

	l, m, n := 1, 2, 3
	fmt.Println(l, m, n)
}
