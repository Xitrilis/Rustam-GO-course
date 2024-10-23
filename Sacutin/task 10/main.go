package main

import "fmt"

func main() {
	var one, two, three int

	fmt.Print("Введите 1-ое число:")
	fmt.Scan(&one)
	fmt.Print("Введите 2-ое число:")
	fmt.Scan(&two)
	fmt.Print("Введите 3-ое число:")
	fmt.Scan(&three)

	if one == two || two == three || three == one {
		fmt.Println("1 число =", one+5)
		fmt.Println("2 число =", two+5)
		fmt.Println("3 число =", three+5)
	} else {
		fmt.Println("Равных нет")
	}
}
