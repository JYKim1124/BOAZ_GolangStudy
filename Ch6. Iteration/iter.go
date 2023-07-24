package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1부터 10까지 정수 합계 : ", sum)

	n := 2

	for n < 100 {
		fmt.Printf("count %d\n", n)

		n *= 2
	}

	for {
		fmt.Printf("무한루프입니다.\n")
	}

	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}

	for index, num := range arr {
		fmt.Printf("arr[%d]의 값은 %d입니다.\n", index, num)
	}

	var arr2 [4]int = [4]int{7, 8, 9, 10}

	for _, n := range arr2 {
		fmt.Printf("%d\n", n)
	}

	var fruits map[string]string = map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"grape":  "purple",
	}

	for fruit, color := range fruits {
		fmt.Printf("%s의 색깔은 %s입니다.", fruit, color)
	}
}
