package main

import "fmt"

func main() {
	fmt.Println(MultiplesNM(1000, 3, 5))
}

// MultiplesNM - Кратные N и M и сложение
func MultiplesNM(inp, n, m int) int {
	var sum int

	for i := 0; i < inp; i++ {
		if i%n == 0 {
			fmt.Println(i)
			sum += i
		}
		if i%m == 0 {
			fmt.Println(i)
			sum += i
		}
	}
	return sum
}
