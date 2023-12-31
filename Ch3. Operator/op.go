package main

import "fmt"

func main() {
	num1, num2 := 17, 5
	str1, str2 := "Hello", "goorm!"

	fmt.Println("num1 + num2 =", num1+num2)
	fmt.Println("str1 + str2 =", str1+str2)
	fmt.Println("num1 - num2 =", num1-num2)
	fmt.Println("num1 * num2 =", num1*num2)
	fmt.Println("num1 / num2 =", num1/num2)
	fmt.Println("num1 % num2 =", num1%num2)

	count1, count2 := 1, 10.4

	count1++
	count2--

	fmt.Println("count1++ :", count1)
	fmt.Println("count2-- :", count2)

	a := 2
	var num int

	num = a
	fmt.Println("num = a :", num)
	num += 4
	fmt.Println("num += 4 :", num)
	num -= 2
	fmt.Println("num -= 2 :", num)
	num *= 5
	fmt.Println("num *= 5 :", num)
	num /= 2
	fmt.Println("num /= 2 :", num)
	num %= 3
	fmt.Println("num %= 3 :", num)

	num = 3  //00000011
	num &= 2 //00000010
	fmt.Printf("num &= 2 : %08b, %d\n", num, num)
	num |= 5 //00000101
	fmt.Printf("num |= 5 : %08b, %d\n", num, num)
	num ^= 4 //00000100
	fmt.Printf("num ^= 4 : %08b, %d\n", num, num)
	num &^= 2 //00000010
	fmt.Printf("num &^= 2 : %08b, %d\n", num, num)
	num <<= 9 //00001001
	fmt.Printf("num <<= 9 : %08b, %d\n", num, num)
	num >>= 8 //00001000
	fmt.Printf("num >>= 8 : %08b, %d\n", num, num)

	fmt.Printf("num1 &^ num2 : %08b, %d\n", num1&^num2, num1&^num2) // And Not 연산
}
