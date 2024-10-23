package main

import (
	"fmt"
)

func main() {
	var one int
	var two int
	fmt.Print("Введите 1 число:")
	fmt.Scan(&one)
	fmt.Print("Введите 2 число:")
	fmt.Scan(&two)

	x := one - two
	y := two - one

	if one > two {
		fmt.Println("Первое число больше второго на", x)
	} else if one == two {
		fmt.Println("Оба числа равны")
	} else {
		fmt.Println("Второе число больше первого на", y)
	}
}
