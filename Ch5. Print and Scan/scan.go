package main

import "fmt"

func main() {
	var i, j int

	fmt.Print("주민등록번호를 -를 이용해 입력하세요 :")
	fmt.Scanf("%d-%d", &i, &j)

	fmt.Printf("주민등록번호는 %d-%d\n", i, j)
	fmt.Println("앞자리는", i)
	fmt.Println("뒷자리는", j)
}
