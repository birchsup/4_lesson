package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введите длинну массива:")
	var n int
	var b int
	fmt.Scanln(&n)
	mas := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		{
			fmt.Println("Введите ", i+1, " число:")
			fmt.Scanln(&b)
			mas = append(mas, float64(b))
		}
		for i := 1; i < len(mas); i++ {
			x := mas[i]
			j := i
			for ; j >= 1 && mas[j-1] > x; j-- {
				mas[j] = mas[j-1]
			}
			mas[j] = x
		}
	}

	fmt.Println(mas)
}
