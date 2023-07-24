package main

import "fmt"

func main() {
	var num int

	fmt.Print("정수입력 :")
	fmt.Scan(&num)

	if num == 1 {
		fmt.Print("hello\n")
	} else if num == 2 {
		fmt.Print("world\n")
	} else {
		fmt.Print("worng number..\n")
	}

	var num2 int

	fmt.Print("정수입력 : ")
	fmt.Scan(&num2)

	if val := num2 * 2; val == 2 {
		fmt.Print("hello\n")
	} else if val := num2 * 3; val == 6 {
		fmt.Print("world\n")
	} else {
		fmt.Print("worng number..\n")
	}
	var opt int
	var num3, num4, result float32

	fmt.Print("1.덧셈 2.뺄셈 3.곱셈 4.나눗셈 선택:")
	fmt.Scan(&opt)
	fmt.Print("두 개의 실수 입력:")
	fmt.Scan(&num3, &num4)

	if opt == 1 {
		result = num3 + num4
	} else if opt == 2 {
		result = num3 - num4
	} else if opt == 3 {
		result = num3 * num4
	} else if opt == 4 {
		result = num3 / num4
	}

	fmt.Printf("결과: %f\n", result)
}
