package main

import "fmt"

func main() {
	var num int = 10
	var changef float32 = float32(num)
	changei := int8(num)

	var str string = "goorm"
	changestr := []byte(str)  //바이트 배열
	str2 := string(changestr) //바이트 배열을 다시 문자열로 변환

	fmt.Println(num)
	fmt.Println(changef, changei)

	fmt.Println(str)
	fmt.Println(changestr)
	fmt.Println(str2)
}
